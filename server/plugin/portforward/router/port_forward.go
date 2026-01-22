package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var PortForward = new(portForward)

type portForward struct{}

// Init 初始化 端口转发 路由信息
func (r *portForward) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("portForward").Use(middleware.OperationRecord())
		group.POST("createPortForward", apiPortForward.CreatePortForward)             // 新建端口转发规则
		group.DELETE("deletePortForward", apiPortForward.DeletePortForward)           // 删除端口转发规则
		group.DELETE("deletePortForwardByIds", apiPortForward.DeletePortForwardByIds) // 批量删除端口转发规则
		group.PUT("updatePortForward", apiPortForward.UpdatePortForward)              // 更新端口转发规则
		group.PUT("updatePortForwardStatus", apiPortForward.UpdatePortForwardStatus)  // 更新端口转发规则状态
	}
	{
		group := private.Group("portForward")
		group.GET("findPortForward", apiPortForward.FindPortForward)           // 根据ID获取端口转发规则
		group.GET("getPortForwardList", apiPortForward.GetPortForwardList)     // 获取端口转发规则列表
		group.GET("getServerIP", apiPortForward.GetServerIP)                   // 获取服务器IP地址
		group.GET("getForwarderStatus", apiPortForward.GetForwarderStatus)     // 获取端口转发运行状态
		group.GET("getAllForwarderStatus", apiPortForward.GetAllForwarderStatus) // 获取所有端口转发运行状态
	}
}
