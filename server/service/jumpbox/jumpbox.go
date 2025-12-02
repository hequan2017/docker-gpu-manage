package jumpbox

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net"
	"os"
	"path/filepath"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
)

// JumpboxService SSH跳板机服务
type JumpboxService struct {
	config *ssh.ServerConfig
	port   int
}

var jumpboxService *JumpboxService

// StartJumpboxServer 启动SSH跳板机服务器
func StartJumpboxServer() error {
	if !global.GVA_CONFIG.Jumpbox.Enabled {
		global.GVA_LOG.Info("SSH跳板机未启用")
		return nil
	}

	port := global.GVA_CONFIG.Jumpbox.Port
	if port == 0 {
		port = 2026 // 默认端口
	}

	// 生成或加载主机密钥
	signer, err := getHostKey()
	if err != nil {
		return fmt.Errorf("获取主机密钥失败: %v", err)
	}

	// 配置SSH服务器
	config := &ssh.ServerConfig{
		PasswordCallback: PasswordAuth,
	}
	config.AddHostKey(signer)

	jumpboxService = &JumpboxService{
		config: config,
		port:   port,
	}

	// 启动服务器
	go jumpboxService.listen()

	global.GVA_LOG.Info("SSH跳板机服务器已启动", zap.Int("port", port))
	return nil
}

// listen 监听连接
func (s *JumpboxService) listen() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		global.GVA_LOG.Error("SSH跳板机监听失败", zap.Error(err), zap.Int("port", s.port))
		return
	}
	defer listener.Close()

	global.GVA_LOG.Info("SSH跳板机正在监听", zap.String("address", listener.Addr().String()))

	for {
		conn, err := listener.Accept()
		if err != nil {
			global.GVA_LOG.Error("接受SSH连接失败", zap.Error(err))
			continue
		}

		// 为每个连接启动goroutine
		go s.handleConnection(conn)
	}
}

// handleConnection 处理SSH连接
func (s *JumpboxService) handleConnection(conn net.Conn) {
	defer conn.Close()

	// 升级为SSH连接
	sshConn, chans, reqs, err := ssh.NewServerConn(conn, s.config)
	if err != nil {
		global.GVA_LOG.Warn("SSH握手失败", zap.Error(err), zap.String("remote", conn.RemoteAddr().String()))
		return
	}
	defer sshConn.Close()

	global.GVA_LOG.Info("SSH连接建立",
		zap.String("username", sshConn.User()),
		zap.String("remote", conn.RemoteAddr().String()),
	)

	// 获取用户信息
	userID, authorityID, username := GetUserFromPermissions(sshConn.Permissions)
	global.GVA_LOG.Info("SSH用户认证成功",
		zap.String("username", username),
		zap.Uint("userId", userID),
		zap.Uint("authorityId", authorityID),
	)

	// 处理全局请求
	go func() {
		for req := range reqs {
			switch req.Type {
			case "keepalive@openssh.com":
				req.Reply(true, nil)
			default:
				req.Reply(false, nil)
			}
		}
	}()

	// 处理通道请求
	for newChannel := range chans {
		go HandleSession(newChannel, userID, authorityID)
	}
	
	// 连接会保持打开直到客户端断开
	// chans通道关闭时，循环退出，函数返回，连接关闭
}

// getHostKey 获取或生成主机密钥
func getHostKey() (ssh.Signer, error) {
	// 如果配置了密钥路径，尝试加载
	if global.GVA_CONFIG.Jumpbox.HostKey != "" {
		keyPath := global.GVA_CONFIG.Jumpbox.HostKey
		if _, err := os.Stat(keyPath); err == nil {
			privateBytes, err := os.ReadFile(keyPath)
			if err != nil {
				return nil, fmt.Errorf("读取密钥文件失败: %v", err)
			}

			signer, err := ssh.ParsePrivateKey(privateBytes)
			if err != nil {
				return nil, fmt.Errorf("解析密钥失败: %v", err)
			}

			global.GVA_LOG.Info("已加载SSH主机密钥", zap.String("path", keyPath))
			return signer, nil
		}
	}

	// 生成新的密钥
	global.GVA_LOG.Info("生成新的SSH主机密钥")
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, fmt.Errorf("生成RSA密钥失败: %v", err)
	}

	// 如果配置了密钥路径，保存密钥
	if global.GVA_CONFIG.Jumpbox.HostKey != "" {
		keyPath := global.GVA_CONFIG.Jumpbox.HostKey
		keyDir := filepath.Dir(keyPath)
		if err := os.MkdirAll(keyDir, 0700); err != nil {
			global.GVA_LOG.Warn("创建密钥目录失败", zap.Error(err))
		} else {
			// 保存私钥
			privateKeyPEM := &pem.Block{
				Type:  "RSA PRIVATE KEY",
				Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
			}
			privateKeyBytes := pem.EncodeToMemory(privateKeyPEM)
			if err := os.WriteFile(keyPath, privateKeyBytes, 0600); err != nil {
				global.GVA_LOG.Warn("保存密钥文件失败", zap.Error(err))
			} else {
				global.GVA_LOG.Info("已保存SSH主机密钥", zap.String("path", keyPath))
			}
		}
	}

	// 转换为SSH签名器
	signer, err := ssh.NewSignerFromKey(privateKey)
	if err != nil {
		return nil, fmt.Errorf("创建签名器失败: %v", err)
	}

	return signer, nil
}

// getBanner 获取SSH欢迎信息
func getBanner() string {
	banner := global.GVA_CONFIG.Jumpbox.Banner
	if banner == "" {
		banner = "欢迎使用SSH跳板机服务\r\n"
	}
	return banner
}

