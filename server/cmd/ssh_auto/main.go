package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

// 自动连接 SSH，匹配包含 targetID 的容器行，输入序号后执行 date。
func main() {
	const (
		user     = "admin"
		addr     = "192.168.112.148:2026"
		targetID = "c61c436b3ef7"
	)

	password := os.Getenv("SSH_PASSWORD")
	if password == "" {
		log.Fatal("未设置 SSH_PASSWORD 环境变量")
	}

	cfg := &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 如需严格校验可替换
		Timeout:         10 * time.Second,
	}

	client, err := ssh.Dial("tcp", addr, cfg)
	if err != nil {
		log.Fatalf("连接 SSH 失败: %v", err)
	}
	defer client.Close()

	sess, err := client.NewSession()
	if err != nil {
		log.Fatalf("创建 session 失败: %v", err)
	}
	defer sess.Close()

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	if err := sess.RequestPty("xterm", 80, 40, modes); err != nil {
		log.Fatalf("请求 PTY 失败: %v", err)
	}

	stdin, err := sess.StdinPipe()
	if err != nil {
		log.Fatalf("获取 stdin 失败: %v", err)
	}
	stdout, err := sess.StdoutPipe()
	if err != nil {
		log.Fatalf("获取 stdout 失败: %v", err)
	}
	stderr, err := sess.StderrPipe()
	if err != nil {
		log.Fatalf("获取 stderr 失败: %v", err)
	}

	if err := sess.Shell(); err != nil {
		log.Fatalf("启动 shell 失败: %v", err)
	}

	combined := io.MultiReader(stdout, stderr)
	scanner := bufio.NewScanner(combined)
	scanner.Buffer(make([]byte, 0, 64*1024), 1*1024*1024)

	var buf bytes.Buffer
	targetIdx := ""
	matchedCh := make(chan struct{}, 1)
	doneCh := make(chan struct{})

	go func() {
		defer close(doneCh)
		for scanner.Scan() {
			line := scanner.Text()
			buf.WriteString(line + "\n")

			if targetIdx == "" && strings.Contains(line, targetID) {
				fields := strings.Fields(line)
				if len(fields) > 0 {
					targetIdx = fields[0]
					fmt.Fprintf(stdin, "%s\n", targetIdx)
					fmt.Fprintf(stdin, "date\n")
					matchedCh <- struct{}{}
				}
			}
		}
	}()

	// 等待匹配及输出
	select {
	case <-matchedCh:
		time.Sleep(3 * time.Second) // 给 date 输出时间
	case <-time.After(15 * time.Second):
		buf.WriteString("等待匹配超时\n")
	}

	_ = sess.Close()
	<-doneCh

	fmt.Println("=== 捕获输出 ===")
	fmt.Println(buf.String())

	if targetIdx == "" {
		log.Fatalf("未匹配到包含 %s 的容器行", targetID)
	}
}
