package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// 简单的Echo服务器 - 用于测试端口转发
// 监听在指定端口，将接收到的数据原样返回
func main() {
	// 配置
	port := "9999" // 默认监听端口

	// 如果提供了命令行参数，使用第一个参数作为端口
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	// 创建监听地址
	address := ":" + port

	// 启动TCP监听
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("❌ 无法启动服务器: %v", err)
	}

	fmt.Printf("✅ Echo服务器已启动！\n")
	fmt.Printf("📡 监听地址: %s\n", address)
	fmt.Printf("⏰ 启动时间: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("📝 服务器功能: 接收消息并原样返回\n")
	fmt.Println(strings.Repeat("-", 50))

	// 接受客户端连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("⚠️  接受连接失败: %v", err)
			continue
		}

		// 处理客户端连接
		go handleConnection(conn)
	}
}

// 处理单个客户端连接
func handleConnection(conn net.Conn) {
	defer conn.Close()

	// 获取客户端地址
	remoteAddr := conn.RemoteAddr().String()
	fmt.Printf("\n🔗 新连接来自: %s\n", remoteAddr)

	// 创建读写器
	reader := bufio.NewReader(conn)

	// 循环读取客户端消息
	for {
		// 读取消息（以换行符分隔）
		message, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Printf("⚠️  读取消息错误: %v\n", err)
			}
			break
		}

		// 去除首尾空白
		message = strings.TrimSpace(message)

		// 打印接收到的消息
		fmt.Printf("📨 收到消息: %s\n", message)

		// 构造响应
		response := fmt.Sprintf("Echo: %s [时间: %s]\n",
			message,
			time.Now().Format("15:04:05"))

		// 发送响应
		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Printf("⚠️  发送响应失败: %v\n", err)
			break
		}

		// 打印发送的响应
		fmt.Printf("📤 发送响应: %s", response)
	}

	fmt.Printf("🔌 连接已关闭: %s\n", remoteAddr)
}

func init() {
	// 导入os包
	os.Args = make([]string, 0)
}
