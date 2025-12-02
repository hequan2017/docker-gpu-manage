package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/jumpbox"
)

// InitJumpbox 初始化SSH跳板机服务
func InitJumpbox() {
	if err := jumpbox.StartJumpboxServer(); err != nil {
		panic("启动SSH跳板机服务失败: " + err.Error())
	}
}

