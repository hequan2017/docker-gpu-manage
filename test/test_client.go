package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strings"
	"time"
)

// 简单的TCP客户端 - 用于测试端口转发
// 连接到指定地址和端口，发送消息并接收响应
func main() {
	// 命令行参数
	server := flag.String("server", "127.0.0.1:8081", "服务器地址 (默认: 127.0.0.1:8081)")
	message := flag.String("msg", "Hello Port Forwarding!", "要发送的消息")
	count := flag.Int("count", 1, "发送消息的次数")
	interval := flag.Int("interval", 1000, "发送间隔(毫秒)")
	flag.Parse()

	fmt.Printf("========================================\n")
	fmt.Printf("     端口转发测试客户端\n")
	fmt.Printf("========================================\n")
	fmt.Printf("目标服务器: %s\n", *server)
	fmt.Printf("测试消息: %s\n", *message)
	fmt.Printf("发送次数: %d\n", *count)
	fmt.Printf("发送间隔: %d ms\n", *interval)
	fmt.Printf("========================================\n\n")

	// 执行测试
	successCount := 0
	failCount := 0

	for i := 1; i <= *count; i++ {
		fmt.Printf("【第 %d/%d 次测试】\n", i, *count)

		err := sendTestMessage(*server, *message, i)
		if err != nil {
			fmt.Printf("❌ 测试失败: %v\n\n", err)
			failCount++
		} else {
			successCount++
		}

		// 如果不是最后一次，等待指定间隔
		if i < *count {
			time.Sleep(time.Duration(*interval) * time.Millisecond)
		}
	}

	// 打印测试结果
	fmt.Printf("\n========================================\n")
	fmt.Printf("           测试结果统计\n")
	fmt.Printf("========================================\n")
	fmt.Printf("总测试次数: %d\n", *count)
	fmt.Printf("✅ 成功: %d\n", successCount)
	fmt.Printf("❌ 失败: %d\n", failCount)
	if successCount == *count {
		fmt.Printf("\n🎉 所有测试通过！端口转发工作正常！\n")
	} else {
		fmt.Printf("\n⚠️  部分测试失败，请检查配置\n")
	}
	fmt.Printf("========================================\n")
}

// 发送测试消息
func sendTestMessage(serverAddr, message string, testNum int) error {
	// 建立连接
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		return fmt.Errorf("连接失败: %v", err)
	}
	defer conn.Close()

	fmt.Printf("📡 已连接到服务器: %s\n", serverAddr)
	fmt.Printf("⏰ 时间: %s\n", time.Now().Format("15:04:05.000"))

	// 发送消息
	fullMessage := fmt.Sprintf("[%d] %s\n", testNum, message)
	_, err = conn.Write([]byte(fullMessage))
	if err != nil {
		return fmt.Errorf("发送失败: %v", err)
	}

	fmt.Printf("📤 发送: %s", fullMessage)

	// 设置读取超时
	deadline := time.Now().Add(5 * time.Second)
	err = conn.SetReadDeadline(deadline)
	if err != nil {
		return fmt.Errorf("设置超时失败: %v", err)
	}

	// 接收响应
	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("接收响应失败: %v", err)
	}

	response = strings.TrimSpace(response)
	fmt.Printf("📥 接收: %s\n", response)

	// 验证响应
	if strings.Contains(response, "Echo:") {
		fmt.Printf("✅ 测试通过\n")
		return nil
	} else {
		return fmt.Errorf("响应格式异常")
	}
}
