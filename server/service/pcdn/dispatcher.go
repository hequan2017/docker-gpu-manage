package pcdn

import (
	"context"
	"fmt"
	"time"
)

// DispatchRequest 下发请求。
type DispatchRequest struct {
	TaskID     string
	TraceID    string
	ContentID  string
	TargetNode uint
	TimeoutSec int
}

// DispatchResult 下发结果。
type DispatchResult struct {
	Success bool
	Message string
}

// Dispatcher 抽象下发协议接口，便于后续扩展 HTTP/gRPC/MQ。
type Dispatcher interface {
	Dispatch(ctx context.Context, request DispatchRequest) DispatchResult
	Protocol() string
}

// MockDispatcher 为默认实现。
type MockDispatcher struct{}

func (d *MockDispatcher) Protocol() string { return "mock" }

func (d *MockDispatcher) Dispatch(ctx context.Context, request DispatchRequest) DispatchResult {
	if request.TimeoutSec <= 0 {
		request.TimeoutSec = 8
	}
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Duration(request.TimeoutSec)*time.Second)
	defer cancel()
	select {
	case <-timeoutCtx.Done():
		return DispatchResult{Success: false, Message: timeoutCtx.Err().Error()}
	case <-time.After(50 * time.Millisecond):
		if request.TargetNode == 0 {
			return DispatchResult{Success: false, Message: "invalid target node"}
		}
		return DispatchResult{Success: true, Message: fmt.Sprintf("dispatched to node %d", request.TargetNode)}
	}
}
