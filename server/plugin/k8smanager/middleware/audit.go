package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	// AuditContextKey 审计上下文键
	AuditContextKey = "k8s_audit_context"
)

// AuditContext 审计上下文
type AuditContext struct {
	StartTime   time.Time
	Action      string
	Resource    string
	ClusterName string
	Namespace   string
	Description string
}

// K8sAuditMiddleware K8s审计中间件
func K8sAuditMiddleware() gin.HandlerFunc {
	auditService := &service.K8sAuditService{}

	return func(c *gin.Context) {
		startTime := time.Now()

		// 读取请求体（用于记录）
		var requestBody interface{}
		if c.Request.Body != nil && c.Request.Method != "GET" {
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err == nil {
				// 恢复请求体
				c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
				json.Unmarshal(bodyBytes, &requestBody)
			}
		}

		// 创建审计上下文
		auditCtx := &AuditContext{
			StartTime: startTime,
		}

		// 从路径解析操作类型和资源类型
		path := c.Request.URL.Path
		if strings.HasPrefix(path, "/k8s/") {
			parts := strings.Split(strings.TrimPrefix(path, "/k8s/"), "/")
			if len(parts) >= 2 {
				auditCtx.Resource = parts[0]           // cluster, pod, deployment等
				auditCtx.Action = getActionFromMethod(c.Request.Method, parts[1]) // create, delete, update等
			}
		}

		// 将审计上下文存入 Gin 上下文
		c.Set(AuditContextKey, auditCtx)

		// 使用响应写入器包装以捕获响应
		w := &responseWriter{
			ResponseWriter: c.Writer,
			statusCode:     200,
		}
		c.Writer = w

		// 继续处理请求
		c.Next()

		// 请求处理完成后记录审计日志
		duration := time.Since(startTime)

		// 只记录成功的请求或失败的请求（跳过404等）
		if w.statusCode >= 200 && w.statusCode < 500 {
			// 获取用户信息
			userID, _ := c.Get("userId")
			username, _ := c.Get("username")

			// 获取集群名称和命名空间
			clusterName := c.Query("clusterName")
			if clusterName == "" {
				clusterName = c.PostForm("clusterName")
			}
			if clusterName == "" && auditCtx != nil {
				clusterName = auditCtx.ClusterName
			}

			namespace := c.Query("namespace")
			if namespace == "" {
				namespace = c.PostForm("namespace")
			}

			// 从请求体中提取
			if reqMap, ok := requestBody.(map[string]interface{}); ok {
				if clusterName == "" {
					if cn, ok := reqMap["clusterName"].(string); ok {
						clusterName = cn
					}
				}
				if namespace == "" {
					if ns, ok := reqMap["namespace"].(string); ok {
						namespace = ns
					}
				}
			}

			// 确定操作状态
			status := "success"
			errorMsg := ""
			if len(c.Errors) > 0 {
				status = "failure"
				errorMsg = c.Errors.String()
			} else if w.statusCode >= 400 {
				status = "failure"
				errorMsg = "HTTP status error"
			}

			// 获取用户ID
			var uid uint
			if userID != nil {
				uid, _ = userID.(uint)
			}

			// 获取用户名
			var uname string
			if username != nil {
				uname, _ = username.(string)
			}

			// 生成资源ID
			resourceID := generateResourceID(auditCtx.Resource, requestBody, c.Params)

			record := &service.AuditRecord{
				UserID:      uid,
				Username:    uname,
				UserIP:      c.ClientIP(),
				UserAgent:   c.Request.UserAgent(),
				Action:      auditCtx.Action,
				Resource:    auditCtx.Resource,
				ResourceID:  resourceID,
				ClusterName: clusterName,
				Namespace:   namespace,
				Description: generateDescription(auditCtx.Action, auditCtx.Resource, resourceID),
				RequestData: requestBody,
				Response:    nil, // 不记录完整响应以节省空间
				Status:      status,
				ErrorMsg:    errorMsg,
				Metadata: map[string]interface{}{
					"method":       c.Request.Method,
					"path":         path,
					"status_code":  w.statusCode,
					"duration_ms":  duration.Milliseconds(),
				},
			}

			// 异步记录审计日志
			go func() {
				if err := auditService.LogOperation(c.Request.Context(), record); err != nil {
					global.GVA_LOG.Error("记录K8s审计日志失败", zap.Error(err))
				}
			}()
		}
	}
}

// getActionFromMethod 根据HTTP方法和路径操作确定操作类型
func getActionFromMethod(method, operation string) string {
	switch operation {
	case "create":
		return "create"
	case "delete", "deleteByIds":
		return "delete"
	case "update":
		return "update"
	case "get", "list", "all":
		return "query"
	case "scale":
		return "scale"
	case "restart":
		return "restart"
	case "refresh":
		return "refresh"
	case "log":
		return "view_logs"
	case "events":
		return "view_events"
	case "containers":
		return "query"
	case "endpoints":
		return "query"
	default:
		// 根据HTTP方法推断
		switch method {
		case "POST":
			return "create"
		case "PUT", "PATCH":
			return "update"
		case "DELETE":
			return "delete"
		case "GET":
			return "query"
		default:
			return "unknown"
		}
	}
}

// generateResourceID 生成资源标识
func generateResourceID(resourceType string, requestBody interface{}, params gin.Params) string {
	// 尝试从请求体获取
	if reqMap, ok := requestBody.(map[string]interface{}); ok {
		if name, ok := reqMap["name"].(string); ok {
			return name
		}
		if id, ok := reqMap["id"]; ok {
			return fmt.Sprintf("%v", id)
		}
	}

	// 尝试从URL参数获取
	for _, param := range params {
		if param.Key == "name" || param.Key == "id" {
			return param.Value
		}
	}

	return ""
}

// generateDescription 生成操作描述
func generateDescription(action, resource, resourceID string) string {
	actionText := map[string]string{
		"create":      "创建",
		"delete":      "删除",
		"update":      "更新",
		"query":       "查询",
		"scale":       "扩缩容",
		"restart":     "重启",
		"refresh":     "刷新",
		"view_logs":   "查看日志",
		"view_events": "查看事件",
	}

	text, ok := actionText[action]
	if !ok {
		text = action
	}

	resourceText := map[string]string{
		"cluster":     "集群",
		"pod":         "Pod",
		"deployment":  "Deployment",
		"service":     "Service",
		"namespace":   "命名空间",
		"event":       "事件",
	}

	rText, ok := resourceText[resource]
	if !ok {
		rText = resource
	}

	if resourceID != "" {
		return fmt.Sprintf("%s%s(%s)", text, rText, resourceID)
	}
	return fmt.Sprintf("%s%s", text, rText)
}

// responseWriter 响应写入器，用于捕获状态码
type responseWriter struct {
	gin.ResponseWriter
	statusCode int
}

func (w *responseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *responseWriter) Write(b []byte) (int, error) {
	return w.ResponseWriter.Write(b)
}

func (w *responseWriter) WriteString(s string) (int, error) {
	return w.ResponseWriter.WriteString(s)
}
