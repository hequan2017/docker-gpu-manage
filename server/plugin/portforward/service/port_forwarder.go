package service

import (
	"context"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/portforward/model"
	"go.uber.org/zap"
)

var (
	forwarderManager *ForwarderManager
	forwarderOnce    sync.Once
)

// ForwarderManager 端口转发管理器
type ForwarderManager struct {
	forwarders map[uint]*Forwarder // ID -> 转发器
	mu         sync.RWMutex        // 读写锁
	logger     *zap.Logger         // 日志记录器
	ctx        context.Context     // 上下文
	cancel     context.CancelFunc  // 取消函数
}

// Forwarder 端口转发器
type Forwarder struct {
	ID        uint                  // 规则ID
	Rule      *model.PortForward    // 转发规则
	Listener  net.Listener          // TCP监听器
	Conn      *net.UDPConn          // UDP连接
	Active    bool                  // 是否活跃
	Cancel    context.CancelFunc    // 取消函数
	mu        sync.Mutex            // 锁
	logger    *zap.Logger           // 日志记录器
	ConnCount int                   // 活跃连接数
}

// GetForwarderManager 获取转发管理器单例
func GetForwarderManager(logger *zap.Logger) *ForwarderManager {
	forwarderOnce.Do(func() {
		ctx, cancel := context.WithCancel(context.Background())
		forwarderManager = &ForwarderManager{
			forwarders: make(map[uint]*Forwarder),
			logger:     logger,
			ctx:        ctx,
			cancel:     cancel,
		}
	})
	return forwarderManager
}

// StartPortForward 启动端口转发
func (fm *ForwarderManager) StartPortForward(rule *model.PortForward) error {
	fm.mu.Lock()
	defer fm.mu.Unlock()

	// 检查是否已经存在转发器
	if _, exists := fm.forwarders[rule.ID]; exists {
		return fmt.Errorf("端口转发规则 %d 已在运行中", rule.ID)
	}

	// 创建转发器
	forwarder := &Forwarder{
		ID:     rule.ID,
		Rule:   rule,
		Active: false,
		logger: fm.logger,
	}

	// 根据协议类型启动转发
	switch rule.Protocol {
	case "tcp":
		if err := fm.startTCPForwarder(forwarder); err != nil {
			return err
		}
	case "udp":
		if err := fm.startUDPForwarder(forwarder); err != nil {
			return err
		}
	default:
		return fmt.Errorf("不支持的协议类型: %s", rule.Protocol)
	}

	// 保存转发器
	fm.forwarders[rule.ID] = forwarder

	fm.logger.Info("端口转发已启动",
		zap.Uint("rule_id", rule.ID),
		zap.String("source", fmt.Sprintf("%s:%d", rule.SourceIP, rule.SourcePort)),
		zap.String("target", fmt.Sprintf("%s:%d", rule.TargetIP, rule.TargetPort)),
		zap.String("protocol", rule.Protocol),
	)

	return nil
}

// startTCPForwarder 启动TCP转发器
func (fm *ForwarderManager) startTCPForwarder(forwarder *Forwarder) error {
	// 创建监听地址
	listenAddr := fmt.Sprintf("%s:%d", forwarder.Rule.SourceIP, forwarder.Rule.SourcePort)

	// 开始监听
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return fmt.Errorf("TCP监听失败: %v", err)
	}

	forwarder.Listener = listener
	forwarder.Active = true

	// 创建上下文
	ctx, cancel := context.WithCancel(fm.ctx)
	forwarder.Cancel = cancel

	// 启动转发协程
	go fm.handleTCPForwarder(ctx, forwarder)

	return nil
}

// startUDPForwarder 启动UDP转发器
func (fm *ForwarderManager) startUDPForwarder(forwarder *Forwarder) error {
	// 创建监听地址
	listenAddr := fmt.Sprintf("%s:%d", forwarder.Rule.SourceIP, forwarder.Rule.SourcePort)

	// 开始监听
	conn, err := net.ListenPacket("udp", listenAddr)
	if err != nil {
		return fmt.Errorf("UDP监听失败: %v", err)
	}

	forwarder.Conn = conn.(*net.UDPConn)
	forwarder.Active = true

	// 创建上下文
	ctx, cancel := context.WithCancel(fm.ctx)
	forwarder.Cancel = cancel

	// 启动转发协程
	go fm.handleUDPForwarder(ctx, forwarder)

	return nil
}

// handleTCPForwarder 处理TCP转发
func (fm *ForwarderManager) handleTCPForwarder(ctx context.Context, forwarder *Forwarder) {
	defer func() {
		if err := forwarder.Listener.Close(); err != nil {
			fm.logger.Warn("关闭TCP监听器失败",
				zap.Uint("rule_id", forwarder.ID),
				zap.Error(err),
			)
		}
		forwarder.Active = false
	}()

	targetAddr := fmt.Sprintf("%s:%d", forwarder.Rule.TargetIP, forwarder.Rule.TargetPort)

	for {
		select {
		case <-ctx.Done():
			fm.logger.Info("TCP转发器收到停止信号",
				zap.Uint("rule_id", forwarder.ID),
			)
			return
		default:
			// 设置接受超时，避免永久阻塞
			forwarder.Listener.(*net.TCPListener).SetDeadline(time.Now().Add(1 * time.Second))

			// 接受客户端连接
			clientConn, err := forwarder.Listener.Accept()
			if err != nil {
				// 超时错误是正常的，继续循环
				if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
					continue
				}
				// 其他错误可能是停止信号导致的
				if ctx.Err() != nil {
					return
				}
				fm.logger.Error("接受连接失败",
					zap.Uint("rule_id", forwarder.ID),
					zap.Error(err),
				)
				continue
			}

			// 增加连接计数
			forwarder.mu.Lock()
			forwarder.ConnCount++
			forwarder.mu.Unlock()

			// 处理单个连接
			go fm.handleTCPConnection(ctx, forwarder, clientConn, targetAddr)
		}
	}
}

// handleTCPConnection 处理单个TCP连接
func (fm *ForwarderManager) handleTCPConnection(ctx context.Context, forwarder *Forwarder, clientConn net.Conn, targetAddr string) {
	defer func() {
		if err := clientConn.Close(); err != nil {
			fm.logger.Debug("关闭客户端连接失败",
				zap.Uint("rule_id", forwarder.ID),
				zap.Error(err),
			)
		}
		forwarder.mu.Lock()
		forwarder.ConnCount--
		forwarder.mu.Unlock()
	}()

	// 连接到目标服务器
	targetConn, err := net.DialTimeout("tcp", targetAddr, 10*time.Second)
	if err != nil {
		fm.logger.Error("连接目标服务器失败",
			zap.Uint("rule_id", forwarder.ID),
			zap.String("target", targetAddr),
			zap.Error(err),
		)
		return
	}
	defer targetConn.Close()

	fm.logger.Debug("TCP连接已建立",
		zap.Uint("rule_id", forwarder.ID),
		zap.String("client", clientConn.RemoteAddr().String()),
		zap.String("target", targetAddr),
	)

	// 双向转发数据
	go func() {
		defer targetConn.Close()
		defer clientConn.Close()
		_, err := io.Copy(targetConn, clientConn)
		if err != nil {
			fm.logger.Debug("客户端到目标服务器的数据转发错误",
				zap.Uint("rule_id", forwarder.ID),
				zap.Error(err),
			)
		}
	}()

	// 目标服务器到客户端
	_, err = io.Copy(clientConn, targetConn)
	if err != nil {
		fm.logger.Debug("目标服务器到客户端的数据转发错误",
			zap.Uint("rule_id", forwarder.ID),
			zap.Error(err),
		)
	}
}

// handleUDPForwarder 处理UDP转发
func (fm *ForwarderManager) handleUDPForwarder(ctx context.Context, forwarder *Forwarder) {
	defer func() {
		if err := forwarder.Conn.Close(); err != nil {
			fm.logger.Warn("关闭UDP连接失败",
				zap.Uint("rule_id", forwarder.ID),
				zap.Error(err),
			)
		}
		forwarder.Active = false
	}()

	targetAddr := fmt.Sprintf("%s:%d", forwarder.Rule.TargetIP, forwarder.Rule.TargetPort)

	buf := make([]byte, 65535)
	for {
		select {
		case <-ctx.Done():
			fm.logger.Info("UDP转发器收到停止信号",
				zap.Uint("rule_id", forwarder.ID),
			)
			return
		default:
			// 设置读取超时
			forwarder.Conn.SetReadDeadline(time.Now().Add(1 * time.Second))

			// 读取UDP数据包
			n, clientAddr, err := forwarder.Conn.ReadFromUDP(buf)
			if err != nil {
				// 超时错误是正常的
				if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
					continue
				}
				if ctx.Err() != nil {
					return
				}
				fm.logger.Error("读取UDP数据失败",
					zap.Uint("rule_id", forwarder.ID),
					zap.Error(err),
				)
				continue
			}

			// 转发数据包到目标服务器
			go func(data []byte, clientAddr *net.UDPAddr) {
				// 连接到目标服务器
				targetConn, err := net.Dial("udp", targetAddr)
				if err != nil {
					fm.logger.Error("连接目标UDP服务器失败",
						zap.Uint("rule_id", forwarder.ID),
						zap.String("target", targetAddr),
						zap.Error(err),
					)
					return
				}
				defer targetConn.Close()

				// 发送数据到目标
				_, err = targetConn.Write(data)
				if err != nil {
					fm.logger.Error("发送UDP数据到目标失败",
						zap.Uint("rule_id", forwarder.ID),
						zap.Error(err),
					)
					return
				}

				// 读取响应
				response := make([]byte, 65535)
				targetConn.SetReadDeadline(time.Now().Add(5 * time.Second))
				rn, err := targetConn.Read(response)
				if err != nil {
					fm.logger.Debug("读取UDP响应失败",
						zap.Uint("rule_id", forwarder.ID),
						zap.Error(err),
					)
					return
				}

				// 将响应发送回客户端
				_, err = forwarder.Conn.WriteToUDP(response[:rn], clientAddr)
				if err != nil {
					fm.logger.Error("发送UDP响应到客户端失败",
						zap.Uint("rule_id", forwarder.ID),
						zap.Error(err),
					)
				}
			}(buf[:n], clientAddr)
		}
	}
}

// StopPortForward 停止端口转发
func (fm *ForwarderManager) StopPortForward(ruleID uint) error {
	fm.mu.Lock()
	defer fm.mu.Unlock()

	forwarder, exists := fm.forwarders[ruleID]
	if !exists {
		return fmt.Errorf("端口转发规则 %d 不存在或未运行", ruleID)
	}

	// 取消转发器上下文
	if forwarder.Cancel != nil {
		forwarder.Cancel()
	}

	// 关闭监听器
	if forwarder.Listener != nil {
		forwarder.Listener.Close()
	}

	// 关闭UDP连接
	if forwarder.Conn != nil {
		forwarder.Conn.Close()
	}

	// 从map中删除
	delete(fm.forwarders, ruleID)

	fm.logger.Info("端口转发已停止",
		zap.Uint("rule_id", ruleID),
	)

	return nil
}

// IsRunning 检查端口转发是否正在运行
func (fm *ForwarderManager) IsRunning(ruleID uint) bool {
	fm.mu.RLock()
	defer fm.mu.RUnlock()

	forwarder, exists := fm.forwarders[ruleID]
	if !exists {
		return false
	}

	return forwarder.Active
}

// GetRunningForwarders 获取所有运行中的转发器ID列表
func (fm *ForwarderManager) GetRunningForwarders() []uint {
	fm.mu.RLock()
	defer fm.mu.RUnlock()

	var runningIDs []uint
	for id, forwarder := range fm.forwarders {
		if forwarder.Active {
			runningIDs = append(runningIDs, id)
		}
	}

	return runningIDs
}

// GetForwarderStatus 获取转发器状态信息
func (fm *ForwarderManager) GetForwarderStatus(ruleID uint) map[string]interface{} {
	fm.mu.RLock()
	defer fm.mu.RUnlock()

	forwarder, exists := fm.forwarders[ruleID]
	if !exists {
		return map[string]interface{}{
			"running": false,
		}
	}

	return map[string]interface{}{
		"running":    forwarder.Active,
		"protocol":   forwarder.Rule.Protocol,
		"conn_count": forwarder.ConnCount,
	}
}

// StopAll 停止所有端口转发
func (fm *ForwarderManager) StopAll() {
	fm.mu.Lock()
	defer fm.mu.Unlock()

	for _, forwarder := range fm.forwarders {
		if forwarder.Cancel != nil {
			forwarder.Cancel()
		}
		if forwarder.Listener != nil {
			forwarder.Listener.Close()
		}
		if forwarder.Conn != nil {
			forwarder.Conn.Close()
		}
	}

	// 清空map
	fm.forwarders = make(map[uint]*Forwarder)

	fm.logger.Info("所有端口转发已停止")
}

// SyncPortForwardStatus 同步端口转发状态
func (fm *ForwarderManager) SyncPortForwardStatus(rules []model.PortForward) {
	for _, rule := range rules {
		if rule.Status {
			// 规则启用，检查是否在运行
			if !fm.IsRunning(rule.ID) {
				// 未运行，启动它
				if err := fm.StartPortForward(&rule); err != nil {
					fm.logger.Error("启动端口转发失败",
						zap.Uint("rule_id", rule.ID),
						zap.Error(err),
					)
				}
			}
		} else {
			// 规则禁用，检查是否在运行
			if fm.IsRunning(rule.ID) {
				// 正在运行，停止它
				if err := fm.StopPortForward(rule.ID); err != nil {
					fm.logger.Error("停止端口转发失败",
						zap.Uint("rule_id", rule.ID),
						zap.Error(err),
					)
				}
			}
		}
	}

	fm.logger.Info("端口转发状态同步完成",
		zap.Int("total_rules", len(rules)),
		zap.Int("running_forwarders", len(fm.GetRunningForwarders())),
	)
}
