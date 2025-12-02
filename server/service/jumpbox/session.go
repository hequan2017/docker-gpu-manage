package jumpbox

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/computenode"
	instanceModel "github.com/flipped-aurora/gin-vue-admin/server/model/instance"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
)

// InstanceInfo 实例信息
type InstanceInfo struct {
	Index       int
	ID          uint
	Name        string
	ContainerID string
	NodeName    string
}

// HandleSession 处理SSH会话
func HandleSession(newChannel ssh.NewChannel, userID uint, authorityID uint) {
	// 只接受session channel
	if newChannel.ChannelType() != "session" {
		newChannel.Reject(ssh.UnknownChannelType, "unknown channel type")
		return
	}

	channel, requests, err := newChannel.Accept()
	if err != nil {
		global.GVA_LOG.Error("接受SSH通道失败", zap.Error(err))
		return
	}
	defer channel.Close()

	// 处理通道请求
	var shellStarted bool
	shellChan := make(chan bool, 1)
	windowSizeChan := make(chan struct{ width, height uint }, 10) // 缓冲窗口大小变化

	go func() {
		for req := range requests {
			switch req.Type {
			case "pty-req":
				// 接受PTY请求（必须在shell之前）
				req.Reply(true, nil)
			case "shell":
				// 接受shell请求
				if len(req.Payload) > 0 {
					req.Reply(false, nil)
				} else {
					req.Reply(true, nil)
					if !shellStarted {
						shellStarted = true
						// 通知主goroutine开始处理交互
						shellChan <- true
					}
				}
			case "window-change":
				// 处理窗口大小变化，解析并传递
				// SSH window-change请求格式：4字节宽度(大端) + 4字节高度(大端)
				if len(req.Payload) >= 8 {
					width := uint(req.Payload[0])<<24 | uint(req.Payload[1])<<16 | uint(req.Payload[2])<<8 | uint(req.Payload[3])
					height := uint(req.Payload[4])<<24 | uint(req.Payload[5])<<16 | uint(req.Payload[6])<<8 | uint(req.Payload[7])
					select {
					case windowSizeChan <- struct{ width, height uint }{width, height}:
					default:
						// 如果通道满了，丢弃旧值
					}
				}
				req.Reply(true, nil)
			default:
				req.Reply(false, nil)
			}
		}
		close(windowSizeChan)
	}()

	// 等待shell请求，然后处理交互（阻塞直到完成）
	<-shellChan
	handleInteractiveSession(channel, userID, authorityID, windowSizeChan)
}

// handleInteractiveSession 处理交互式会话
func handleInteractiveSession(channel ssh.Channel, userID uint, authorityID uint, windowSizeChan <-chan struct{ width, height uint }) {
	// 判断是否是管理员（authorityID == 888）
	isAdmin := authorityID == 888

	// 获取用户实例列表
	instances, err := getUserInstances(userID, isAdmin)
	if err != nil {
		channel.Write([]byte(fmt.Sprintf("获取实例列表失败: %s\r\n", err.Error())))
		return
	}

	if len(instances) == 0 {
		channel.Write([]byte("您还没有创建任何容器实例。\r\n"))
		channel.Write([]byte("连接将在5秒后关闭...\r\n"))
		time.Sleep(5 * time.Second)
		return
	}

	// 显示容器列表
	channel.Write([]byte("\r\n=== 您的容器列表 ===\r\n"))
	for _, inst := range instances {
		line := fmt.Sprintf("%d  %s-%s-%s\r\n", inst.Index, inst.Name, inst.ContainerID, inst.NodeName)
		channel.Write([]byte(line))
	}

	// 循环处理用户输入，直到输入有效序号或退出
	for {
		// 提示用户输入
		channel.Write([]byte("\r\n请选择要连接的容器（输入序号，输入 'q' 退出）: "))

		// 读取用户输入
		var inputBuf []byte
		buf := make([]byte, 256)

		for {
			n, err := channel.Read(buf)
			if err != nil {
				if err != io.EOF {
					global.GVA_LOG.Error("读取用户输入失败", zap.Error(err))
				}
				return
			}
			if n == 0 {
				continue
			}

			// 查找换行符
			lineEnd := -1
			for i := 0; i < n; i++ {
				if buf[i] == '\n' || buf[i] == '\r' {
					lineEnd = i
					break
				}
			}

			if lineEnd >= 0 {
				inputBuf = append(inputBuf, buf[:lineEnd]...)
				channel.Write(buf[:lineEnd+1])
				break
			} else {
				inputBuf = append(inputBuf, buf[:n]...)
				channel.Write(buf[:n])
			}
		}

		// 处理用户输入
		input := strings.TrimSpace(string(inputBuf))
		if input == "q" || input == "Q" {
			channel.Write([]byte("再见！\r\n"))
			return
		}

		// 解析序号
		index, err := strconv.Atoi(input)
		if err != nil {
			channel.Write([]byte(fmt.Sprintf("无效的序号: %s\r\n", input)))
			continue
		}

		// 查找对应的实例
		var selectedInstance *InstanceInfo
		for _, inst := range instances {
			if inst.Index == index {
				selectedInstance = &inst
				break
			}
		}

		if selectedInstance == nil {
			channel.Write([]byte(fmt.Sprintf("未找到序号为 %d 的容器\r\n", index)))
			continue
		}

		// 找到有效实例，连接到容器
		channel.Write([]byte(fmt.Sprintf("正在连接到容器: %s...\r\n", selectedInstance.Name)))
		channel.Write([]byte("提示：在vim中使用鼠标中键或Shift+Insert可以粘贴剪贴板内容\r\n"))
		connectToContainer(channel, selectedInstance.ID, windowSizeChan)
		return
	}
}

// getUserInstances 获取用户的实例列表
func getUserInstances(userID uint, isAdmin bool) ([]InstanceInfo, error) {
	type InstanceWithNode struct {
		instanceModel.Instance
		NodeName *string `gorm:"column:node_name"`
	}

	var instances []InstanceWithNode

	// 使用JOIN一次性查询实例和节点信息，避免N+1查询问题
	db := global.GVA_DB.Table("instance").
		Select("instance.*, compute_node.name as node_name").
		Joins("LEFT JOIN compute_node ON instance.node_id = compute_node.id").
		Where("instance.deleted_at IS NULL").
		Where("instance.container_id IS NOT NULL AND instance.container_id != ''")

	// 权限控制：普通用户只能看到自己创建的实例，管理员可以看到所有
	if !isAdmin {
		userIDInt64 := int64(userID)
		db = db.Where("instance.user_id = ?", userIDInt64)
	}

	if err := db.Find(&instances).Error; err != nil {
		return nil, err
	}

	// 构建实例信息列表
	result := make([]InstanceInfo, 0, len(instances))
	for i, inst := range instances {
		info := InstanceInfo{
			Index: i,
			ID:    inst.ID,
		}

		if inst.Name != nil {
			info.Name = *inst.Name
		} else {
			info.Name = fmt.Sprintf("instance-%d", inst.ID)
		}

		if inst.ContainerId != nil {
			// 只显示容器ID的前12位（Docker标准格式）
			containerID := *inst.ContainerId
			if len(containerID) > 12 {
				containerID = containerID[:12]
			}
			info.ContainerID = containerID
		}

		// 使用JOIN查询的节点名称
		if inst.NodeName != nil {
			info.NodeName = *inst.NodeName
		} else if inst.NodeId != nil {
			info.NodeName = fmt.Sprintf("node-%d", *inst.NodeId)
		}

		result = append(result, info)
	}

	return result, nil
}

// connectToContainer 连接到容器
func connectToContainer(channel ssh.Channel, instanceID uint, windowSizeChan <-chan struct{ width, height uint }) {
	ctx := context.Background()

	// 使用JOIN一次性查询实例和节点信息，提高查询速度
	type InstanceWithNode struct {
		instanceModel.Instance
		computenode.ComputeNode
	}

	var data InstanceWithNode
	if err := global.GVA_DB.Table("instance").
		Select("instance.*, compute_node.*").
		Joins("LEFT JOIN compute_node ON instance.node_id = compute_node.id").
		Where("instance.id = ?", instanceID).
		Where("instance.deleted_at IS NULL").
		First(&data).Error; err != nil {
		channel.Write([]byte(fmt.Sprintf("获取实例信息失败: %s\r\n", err.Error())))
		return
	}

	inst := data.Instance
	node := data.ComputeNode

	if inst.ContainerId == nil || *inst.ContainerId == "" {
		channel.Write([]byte("容器ID为空\r\n"))
		return
	}

	if inst.NodeId == nil {
		channel.Write([]byte("实例未关联节点\r\n"))
		return
	}

	// 创建Docker客户端
	cli, err := createDockerClient(&node)
	if err != nil {
		channel.Write([]byte(fmt.Sprintf("创建Docker客户端失败: %s\r\n", err.Error())))
		return
	}
	defer cli.Close()

	// 创建exec实例，设置环境变量以支持vim等工具
	// 只设置TERM，不强制设置语言环境，避免容器中没有对应locale时出现警告
	execConfig := container.ExecOptions{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          true,
		Env: []string{
			"TERM=xterm-256color", // 只设置终端类型，支持256色和现代终端特性
		},
		Cmd: []string{"/bin/bash", "-l"}, // 使用login shell，加载完整环境
	}

	// 创建exec，如果失败则尝试sh
	execResp, err := cli.ContainerExecCreate(ctx, *inst.ContainerId, execConfig)
	if err != nil {
		// 如果bash失败，尝试不带-l参数的bash
		execConfig.Cmd = []string{"/bin/bash"}
		execResp, err = cli.ContainerExecCreate(ctx, *inst.ContainerId, execConfig)
		if err != nil {
			// 最后尝试sh
			execConfig.Cmd = []string{"/bin/sh"}
			execResp, err = cli.ContainerExecCreate(ctx, *inst.ContainerId, execConfig)
			if err != nil {
				channel.Write([]byte(fmt.Sprintf("创建exec失败: %s\r\n", err.Error())))
				return
			}
		}
	}

	// 附加到exec
	attachResp, err := cli.ContainerExecAttach(ctx, execResp.ID, container.ExecStartOptions{
		Tty: true,
	})
	if err != nil {
		channel.Write([]byte(fmt.Sprintf("附加到exec失败: %s\r\n", err.Error())))
		return
	}
	defer attachResp.Close()

	// 处理窗口大小变化
	go func() {
		for size := range windowSizeChan {
			err := cli.ContainerExecResize(ctx, execResp.ID, container.ResizeOptions{
				Height: size.height,
				Width:  size.width,
			})
			if err != nil {
				global.GVA_LOG.Debug("调整容器终端大小失败", zap.Error(err))
			}
		}
	}()

	// 转发数据
	go func() {
		io.Copy(channel, attachResp.Reader)
		channel.Close()
	}()

	// 在连接建立后，发送vim配置提示（可选）
	// 注意：这里不能直接发送到attachResp，因为会干扰正常的shell交互
	// 用户需要在vim中使用 :set mouse=a 来启用鼠标支持

	io.Copy(attachResp.Conn, channel)
}

// createDockerClient 创建Docker客户端
func createDockerClient(node *computenode.ComputeNode) (*client.Client, error) {
	if node.DockerAddress == nil || *node.DockerAddress == "" {
		return nil, fmt.Errorf("节点Docker连接地址为空")
	}

	dockerHost := *node.DockerAddress

	// 检查是否使用TLS
	useTLS := node.UseTls != nil && *node.UseTls

	if useTLS {
		// 使用TLS连接
		if node.CaCert == nil || node.ClientCert == nil || node.ClientKey == nil {
			return nil, fmt.Errorf("TLS证书配置不完整")
		}

		// 创建TLS配置
		tlsConfig, err := createTLSConfig(*node.CaCert, *node.ClientCert, *node.ClientKey)
		if err != nil {
			return nil, fmt.Errorf("创建TLS配置失败: %v", err)
		}

		// 创建HTTP客户端
		httpClient := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig,
			},
			Timeout: 30 * time.Second,
		}

		return client.NewClientWithOpts(
			client.WithHost(dockerHost),
			client.WithHTTPClient(httpClient),
			client.WithAPIVersionNegotiation(),
		)
	}

	// 不使用TLS
	return client.NewClientWithOpts(
		client.WithHost(dockerHost),
		client.WithAPIVersionNegotiation(),
	)
}

// createTLSConfig 创建TLS配置
func createTLSConfig(caCert, clientCert, clientKey string) (*tls.Config, error) {
	// 加载CA证书
	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM([]byte(caCert)) {
		return nil, fmt.Errorf("无法解析CA证书")
	}

	// 加载客户端证书
	cert, err := tls.X509KeyPair([]byte(clientCert), []byte(clientKey))
	if err != nil {
		return nil, fmt.Errorf("加载客户端证书失败: %v", err)
	}

	return &tls.Config{
		RootCAs:      caCertPool,
		Certificates: []tls.Certificate{cert},
	}, nil
}
