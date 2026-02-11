package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

func (a *api) HealthCheck(c *gin.Context) {
	response.OkWithMessage("OpenClaw plugin is running", c)
}
