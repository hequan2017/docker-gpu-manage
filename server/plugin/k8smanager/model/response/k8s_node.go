package response

import "time"

type NodeListResponse struct {
	Name          string            `json:"name"`
	Status        string            `json:"status"` // Ready, NotReady, Unknown
	Unschedulable bool              `json:"unschedulable"`
	Roles         string            `json:"roles"`
	Version       string            `json:"version"`
	InternalIP    string            `json:"internalIP"`
	CPU           string            `json:"cpu"`
	Memory        string            `json:"memory"`
	GPU           string            `json:"gpu"` // 显示字符串，如 "1" 或 "0"
	CreatedAt     time.Time         `json:"createdAt"`
	Labels        map[string]string `json:"labels"`
}
