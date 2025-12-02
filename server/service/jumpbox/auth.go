package jumpbox

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
)

// PasswordAuth 密码认证回调函数
func PasswordAuth(conn ssh.ConnMetadata, password []byte) (*ssh.Permissions, error) {
	username := conn.User()
	passwd := string(password)

	// 查询用户
	var user system.SysUser
	if err := global.GVA_DB.Where("username = ? AND deleted_at IS NULL", username).First(&user).Error; err != nil {
		global.GVA_LOG.Warn("SSH登录失败: 用户不存在", zap.String("username", username))
		return nil, ssh.ErrNoAuth
	}

	// 检查用户是否被冻结
	if user.Enable != 1 {
		global.GVA_LOG.Warn("SSH登录失败: 用户已被冻结", zap.String("username", username))
		return nil, ssh.ErrNoAuth
	}

	// 验证密码
	if !utils.BcryptCheck(passwd, user.Password) {
		global.GVA_LOG.Warn("SSH登录失败: 密码错误", zap.String("username", username))
		return nil, ssh.ErrNoAuth
	}

	global.GVA_LOG.Info("SSH登录成功", zap.String("username", username), zap.Uint("userId", user.ID), zap.Uint("authorityId", user.AuthorityId))

	// 返回权限信息，包含用户ID和角色ID
	return &ssh.Permissions{
		Extensions: map[string]string{
			"userId":      strconv.FormatUint(uint64(user.ID), 10),
			"authorityId": strconv.FormatUint(uint64(user.AuthorityId), 10),
			"username":    username,
		},
	}, nil
}

// GetUserFromPermissions 从SSH权限中获取用户信息
func GetUserFromPermissions(perms *ssh.Permissions) (uint, uint, string) {
	if perms == nil || perms.Extensions == nil {
		return 0, 0, ""
	}

	userIdStr := perms.Extensions["userId"]
	authorityIdStr := perms.Extensions["authorityId"]
	username := perms.Extensions["username"]

	userId := uint(0)
	authorityId := uint(0)

	if userIdStr != "" {
		if id, err := strconv.ParseUint(userIdStr, 10, 32); err == nil {
			userId = uint(id)
		}
	}
	if authorityIdStr != "" {
		if id, err := strconv.ParseUint(authorityIdStr, 10, 32); err == nil {
			authorityId = uint(id)
		}
	}

	return userId, authorityId, username
}
