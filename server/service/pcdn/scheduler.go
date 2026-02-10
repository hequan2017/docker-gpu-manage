package pcdn

import (
	"context"
	"errors"
	"sort"
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	computenodeModel "github.com/flipped-aurora/gin-vue-admin/server/model/computenode"
	pcdnModel "github.com/flipped-aurora/gin-vue-admin/server/model/pcdn"
	"gorm.io/gorm"
)

// ScheduleInput 调度输入。
type ScheduleInput struct {
	TaskID      string
	TraceID     string
	ContentID   string
	UserRegion  string
	UserISP     string
	TopN        int
	MaxRetry    int
	TimeoutSec  int
	Protocol    string
	NodeMetrics map[uint]RealtimeNodeMetric
}

// CandidateNode 调度候选结果。
type CandidateNode struct {
	NodeID uint    `json:"nodeId"`
	Name   string  `json:"name"`
	Region string  `json:"region"`
	ISP    string  `json:"isp"`
	Score  float64 `json:"score"`
}

type SchedulerService struct {
	policy *PolicyEngine
}

var PcdnSchedulerService = NewSchedulerService()

func NewSchedulerService() *SchedulerService {
	return &SchedulerService{policy: NewPolicyEngine(ScoreWeights{})}
}

// Schedule 执行“约束过滤 + 加权评分”，并记录到 PcdnDispatchTask。
func (s *SchedulerService) Schedule(ctx context.Context, in ScheduleInput) ([]CandidateNode, error) {
	if in.TaskID == "" || in.TraceID == "" {
		return nil, errors.New("task_id and trace_id are required")
	}
	if in.ContentID == "" || in.UserRegion == "" {
		return nil, errors.New("content_id and user_region are required")
	}
	if in.TopN <= 0 {
		in.TopN = 3
	}
	if in.MaxRetry <= 0 {
		in.MaxRetry = 3
	}
	if in.TimeoutSec <= 0 {
		in.TimeoutSec = 8
	}
	if in.Protocol == "" {
		in.Protocol = "mock"
	}

	var nodes []computenodeModel.ComputeNode
	if err := global.GVA_DB.WithContext(ctx).Where("deleted_at IS NULL").Find(&nodes).Error; err != nil {
		return nil, err
	}

	candidates := make([]CandidateNode, 0, len(nodes))
	for _, node := range nodes {
		if node.ID == 0 || node.IsOnShelf == nil || !*node.IsOnShelf {
			continue
		}
		m, ok := in.NodeMetrics[node.ID]
		if !ok {
			m = RealtimeNodeMetric{NodeID: node.ID, Online: true, HealthScore: 80}
		}
		if !m.Online || m.PolicyDisabled || m.LoadPercent >= 90 {
			continue
		}
		if node.DockerStatus != nil && *node.DockerStatus == "failed" {
			continue
		}

		regionAffinity := 0.92
		if node.Region != nil && *node.Region == in.UserRegion {
			regionAffinity = 1.08
		}
		ispAffinity := 0.95
		if in.UserISP == "" || m.ISP == "" || m.ISP == in.UserISP {
			ispAffinity = 1.03
		}

		score := s.policy.Score(m) * regionAffinity * ispAffinity
		candidates = append(candidates, CandidateNode{
			NodeID: node.ID,
			Name:   safeString(node.Name),
			Region: safeString(node.Region),
			ISP:    m.ISP,
			Score:  score,
		})
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].Score > candidates[j].Score
	})
	if len(candidates) > in.TopN {
		candidates = candidates[:in.TopN]
	}

	if err := s.upsertTask(ctx, in, candidates); err != nil {
		return nil, err
	}
	return candidates, nil
}

func (s *SchedulerService) upsertTask(ctx context.Context, in ScheduleInput, candidates []CandidateNode) error {
	var task pcdnModel.PcdnDispatchTask
	err := global.GVA_DB.WithContext(ctx).Where("task_id = ? AND trace_id = ?", in.TaskID, in.TraceID).First(&task).Error
	if err == nil {
		return nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	candidateList := make([]any, 0, len(candidates))
	for _, c := range candidates {
		candidateList = append(candidateList, map[string]any{
			"node_id": c.NodeID,
			"score":   c.Score,
			"region":  c.Region,
			"isp":     c.ISP,
		})
	}
	metricsSnapshot := map[string]any{}
	for nodeID, m := range in.NodeMetrics {
		metricsSnapshot[strconv.FormatUint(uint64(nodeID), 10)] = map[string]any{
			"latency_ms":      m.LatencyMS,
			"unit_cost":       m.UnitCost,
			"load_percent":    m.LoadPercent,
			"health_score":    m.HealthScore,
			"online":          m.Online,
			"policy_disabled": m.PolicyDisabled,
			"isp":             m.ISP,
		}
	}

	newTask := pcdnModel.PcdnDispatchTask{
		TaskID:           in.TaskID,
		TraceID:          in.TraceID,
		ContentID:        in.ContentID,
		UserRegion:       in.UserRegion,
		UserISP:          in.UserISP,
		TopN:             in.TopN,
		Status:           pcdnModel.DispatchStatusPending,
		MaxRetry:         in.MaxRetry,
		TimeoutSeconds:   in.TimeoutSec,
		Candidates:       map[string]any{"list": candidateList},
		MetricsSnapshot:  metricsSnapshot,
		DispatchProtocol: in.Protocol,
	}
	if len(candidates) > 0 {
		newTask.PrimaryNodeID = candidates[0].NodeID
		newTask.CurrentNodeID = candidates[0].NodeID
	}
	return global.GVA_DB.WithContext(ctx).Create(&newTask).Error
}

func safeString(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}
