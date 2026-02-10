package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/pcdn"

type PcdnNodeListResponse struct {
	List  []pcdn.PcdnNode `json:"list"`
	Total int64           `json:"total"`
	Page  int             `json:"page"`
	Size  int             `json:"size"`
}

type PcdnResourceListResponse struct {
	List  []pcdn.PcdnResource `json:"list"`
	Total int64               `json:"total"`
	Page  int                 `json:"page"`
	Size  int                 `json:"size"`
}

type PcdnPolicyListResponse struct {
	List  []pcdn.PcdnPolicy `json:"list"`
	Total int64             `json:"total"`
	Page  int               `json:"page"`
	Size  int               `json:"size"`
}

type PcdnDispatchTaskListResponse struct {
	List  []pcdn.PcdnDispatchTask `json:"list"`
	Total int64                   `json:"total"`
	Page  int                     `json:"page"`
	Size  int                     `json:"size"`
}

type PcdnMetricSnapshotListResponse struct {
	List  []pcdn.PcdnMetricSnapshot `json:"list"`
	Total int64                     `json:"total"`
	Page  int                       `json:"page"`
	Size  int                       `json:"size"`
}
