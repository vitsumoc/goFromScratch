package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

type Message struct {
	From    string
	Content string
}

var (
	// 存储所有连接的客户端
	clients = make(map[string]net.Conn)
	// 保护clients的互斥锁
	clientsMu sync.Mutex
	// 消息队列
	messages = make([]Message, 0)
	// 保护messages的互斥锁
	messagesMu sync.Mutex
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8864")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("服务启动，监听：0.0.0.0:8864")

	// 启动消息处理协程
	go processMessages()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println("客户端进入")
		go handleConn(conn)
	}
}

// 处理消息的协程
func processMessages() {
	for {
		time.Sleep(100 * time.Millisecond) // 每100毫秒处理一次消息

		messagesMu.Lock()
		if len(messages) == 0 {
			messagesMu.Unlock()
			continue
		}

		// 获取所有待处理的消息
		msgs := make([]Message, len(messages))
		copy(msgs, messages)
		messages = messages[:0] // 清空消息队列
		messagesMu.Unlock()

		// 发送消息给所有客户端
		clientsMu.Lock()
		for _, msg := range msgs {
			msgStr := fmt.Sprintf("%s: %s\n", msg.From, msg.Content)
			for addr, conn := range clients {
				_, err := conn.Write([]byte(msgStr))
				if err != nil {
					log.Printf("发送消息到 %s 失败: %v\n", addr, err)
					delete(clients, addr)
					conn.Close()
				}
			}
		}
		clientsMu.Unlock()
	}
}

func handleConn(c net.Conn) {
	who := c.RemoteAddr().String()

	// 注册客户端
	clientsMu.Lock()
	clients[who] = c
	clientsMu.Unlock()

	// 清理函数
	defer func() {
		clientsMu.Lock()
		delete(clients, who)
		clientsMu.Unlock()
		c.Close()
		fmt.Println("客户端退出:", who)
	}()

	// 读取客户端消息
	input := bufio.NewScanner(c)
	for input.Scan() {
		// 将消息加入队列
		msg := Message{
			From:    who,
			Content: input.Text(),
		}

		messagesMu.Lock()
		messages = append(messages, msg)
		messagesMu.Unlock()
	}

	if err := input.Err(); err != nil {
		log.Printf("读取客户端 %s 消息失败: %v\n", who, err)
	}
}
