package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("使用方法: tcpClient <主机> <端口>")
		fmt.Println("例如: tcpClient 127.0.0.1 7749")
		os.Exit(1)
	}

	host := os.Args[1]
	port := os.Args[2]
	address := fmt.Sprintf("%s:%s", host, port)

	// 建立连接
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("连接失败: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("已连接到 %s\n", address)

	// 创建一个通道用于通知程序退出
	done := make(chan bool)

	// 启动一个 goroutine 用于接收数据
	go func() {
		reader := bufio.NewReader(conn)
		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					fmt.Println("服务器关闭了连接")
					done <- true
					return
				}
				fmt.Printf("读取错误: %v\n", err)
				done <- true
				return
			}
			fmt.Print(message)
		}
	}()

	// 主循环用于发送数据
	scanner := bufio.NewScanner(os.Stdin)
	for {
		select {
		case <-done:
			return
		default:
			if scanner.Scan() {
				message := scanner.Text()
				_, err := fmt.Fprintf(conn, "%s\n", message)
				if err != nil {
					fmt.Printf("发送错误: %v\n", err)
					return
				}
			}
		}
	}
}
