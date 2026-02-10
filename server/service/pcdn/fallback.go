package pcdn

import pcdnModel "github.com/flipped-aurora/gin-vue-admin/server/model/pcdn"

// NextFallbackNode 主节点失败时选择次优候选。
func NextFallbackNode(task pcdnModel.PcdnDispatchTask) uint {
	raw, ok := task.Candidates["list"]
	if ok {
		if nodeID := findNextNode(raw, task.CurrentNodeID); nodeID != 0 {
			return nodeID
		}
	}
	return findNextNode(task.Candidates, task.CurrentNodeID)
}

func findNextNode(raw any, current uint) uint {
	switch list := raw.(type) {
	case []any:
		for _, item := range list {
			nodeID := parseNodeID(item)
			if nodeID != 0 && nodeID != current {
				return nodeID
			}
		}
	case map[string]any:
		nodeID := parseNodeID(list)
		if nodeID != 0 && nodeID != current {
			return nodeID
		}
	}
	return 0
}

func parseNodeID(v any) uint {
	m, ok := v.(map[string]any)
	if !ok {
		return 0
	}
	raw, ok := m["node_id"]
	if !ok {
		return 0
	}
	switch id := raw.(type) {
	case uint:
		return id
	case int:
		if id > 0 {
			return uint(id)
		}
	case float64:
		if id > 0 {
			return uint(id)
		}
	}
	return 0
}
