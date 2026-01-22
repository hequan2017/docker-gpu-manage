package api

import (
	"fmt"
	"net"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/portforward/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/portforward/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var PortForward = new(portForward)

type portForward struct{}

// CreatePortForward 创建端口转发规则
// @Tags PortForward
// @Summary 创建端口转发规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PortForward true "源IP、源端口、协议、目标IP、目标端口"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /portForward/createPortForward [post]
func (a *portForward) CreatePortForward(c *gin.Context) {
	var portForward model.PortForward
	err := c.ShouldBindJSON(&portForward)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = servicePortForward.CreatePortForward(&portForward)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeletePortForward 删除端口转发规则
// @Tags PortForward
// @Summary 删除端口转发规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PortForward true "删除端口转发规则"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /portForward/deletePortForward [delete]
func (a *portForward) DeletePortForward(c *gin.Context) {
	ID := c.Query("ID")
	err := servicePortForward.DeletePortForward(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeletePortForwardByIds 批量删除端口转发规则
// @Tags PortForward
// @Summary 批量删除端口转发规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /portForward/deletePortForwardByIds [delete]
func (a *portForward) DeletePortForwardByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := servicePortForward.DeletePortForwardByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdatePortForward 更新端口转发规则
// @Tags PortForward
// @Summary 更新端口转发规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PortForward true "更新端口转发规则"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /portForward/updatePortForward [put]
func (a *portForward) UpdatePortForward(c *gin.Context) {
	var portForward model.PortForward
	err := c.ShouldBindJSON(&portForward)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = servicePortForward.UpdatePortForward(portForward)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindPortForward 用id查询端口转发规则
// @Tags PortForward
// @Summary 用id查询端口转发规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PortForward true "用id查询端口转发规则"
// @Success 200 {object} response.Response{data=model.PortForward,msg=string} "查询成功"
// @Router /portForward/findPortForward [get]
func (a *portForward) FindPortForward(c *gin.Context) {
	ID := c.Query("ID")
	rePortForward, err := servicePortForward.GetPortForward(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(rePortForward, c)
}

// GetPortForwardList 分页获取端口转发规则列表
// @Tags PortForward
// @Summary 分页获取端口转发规则列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PortForwardSearch true "分页获取端口转发规则列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /portForward/getPortForwardList [get]
func (a *portForward) GetPortForwardList(c *gin.Context) {
	var pageInfo request.PortForwardSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := servicePortForward.GetPortForwardList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// UpdatePortForwardStatus 更新端口转发规则状态
// @Tags PortForward
// @Summary 更新端口转发规则状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PortForward true "更新端口转发规则状态"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /portForward/updatePortForwardStatus [put]
func (a *portForward) UpdatePortForwardStatus(c *gin.Context) {
	ID := c.Query("ID")
	status := c.Query("status") == "true"
	err := servicePortForward.UpdatePortForwardStatus(ID, status)
	if err != nil {
		global.GVA_LOG.Error("更新状态失败!", zap.Error(err))
		response.FailWithMessage("更新状态失败", c)
		return
	}
	response.OkWithMessage("更新状态成功", c)
}

// GetServerIP 获取服务器IP地址
// @Tags PortForward
// @Summary 获取服务器IP地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /portForward/getServerIP [get]
func (a *portForward) GetServerIP(c *gin.Context) {
	// 获取本机所有网络接口的IP地址
	ips := getLocalIPs()
	response.OkWithData(gin.H{"ips": ips}, c)
}

// getLocalIPs 获取本机所有非127.0.0.1的IP地址
func getLocalIPs() []string {
	var ips []string

	interfaces, err := net.Interfaces()
	if err != nil {
		global.GVA_LOG.Error("获取网络接口失败!", zap.Error(err))
		return ips
	}

	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.IsLoopback() {
				continue
			}

			ip = ip.To4()
			if ip == nil {
				continue
			}

			ips = append(ips, ip.String())
		}
	}

	// 如果没有找到IP，返回默认值
	if len(ips) == 0 {
		ips = append(ips, "0.0.0.0")
	}

	return ips
}

// GetForwarderStatus 获取端口转发状态
// @Tags PortForward
// @Summary 获取端口转发运行状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id query string true "规则ID"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取成功"
// @Router /portForward/getForwarderStatus [get]
func (a *portForward) GetForwarderStatus(c *gin.Context) {
	ID := c.Query("ID")
	if ID == "" {
		response.FailWithMessage("规则ID不能为空", c)
		return
	}

	// 转换ID
	idUint, err := parseUint(ID)
	if err != nil {
		response.FailWithMessage("无效的规则ID", c)
		return
	}

	// 获取转发器状态
	status := servicePortForward.GetForwarderStatus(ID)
	response.OkWithData(status, c)
}

// GetAllForwarderStatus 获取所有端口转发状态
// @Tags PortForward
// @Summary 获取所有端口转发运行状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取成功"
// @Router /portForward/getAllForwarderStatus [get]
func (a *portForward) GetAllForwarderStatus(c *gin.Context) {
	// 获取所有转发器状态
	status := servicePortForward.GetAllForwarderStatus()
	response.OkWithData(status, c)
}

// parseUint 解析uint ID
func parseUint(s string) (uint, error) {
	var id uint64
	_, err := fmt.Sscanf(s, "%d", &id)
	return uint(id), err
}

