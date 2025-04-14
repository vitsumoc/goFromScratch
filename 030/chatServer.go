package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type client struct {
	username    string
	ch          chan string
	enterResult chan bool // 用户名验证结果通道
}

type message struct {
	c       client // 发送方
	content string // 消息内容
}

var (
	clients    = make(map[string]client) // 用户名到client的映射
	entering   = make(chan client)
	leaving    = make(chan client)
	msgChannel = make(chan message) // 所有传入的消息
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:7749")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	log.Println("服务器启动，监听地址：0.0.0.0:7749")

	go handleSig()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("接受连接错误: %v", err)
			continue
		}
		go handleConn(conn)
	}
}

func handleSig() {
	for {
		select {
		case c := <-entering:
			exists := false // 表示用户名是否重复
			for _, client := range clients {
				if client.username == c.username {
					exists = true
					break
				}
			}
			// 未重复则加入聊天室
			if !exists {
				clients[c.username] = c
			}
			c.enterResult <- !exists
		case c := <-leaving:
			delete(clients, c.username)
		case message := <-msgChannel:
			handleMessage(message)
		}
	}
}

func handleConn(conn net.Conn) {
	client := client{
		ch:          make(chan string),
		enterResult: make(chan bool),
	}

	// 启动消息写入协程
	go clientWriter(conn, client)

	// 获取用户名
	fmt.Fprintln(conn, "请输入用户名：")
	input := bufio.NewScanner(conn)
	for input.Scan() {
		username := strings.TrimSpace(input.Text())
		if username == "" {
			fmt.Fprintln(conn, "用户名不能为空，请重新输入：")
			continue
		}
		client.username = username
		// 尝试加入聊天室
		entering <- client
		enterResult := <-client.enterResult
		if !enterResult {
			fmt.Fprintln(conn, "用户名已存在，请重新输入：")
			continue
		}
		break
	}
	// 发送欢迎消息
	client.ch <- fmt.Sprintf("欢迎 %s 加入聊天室！", client.username)
	msgChannel <- message{
		client,
		fmt.Sprintf("%s 加入了聊天室", client.username),
	}

	// 处理用户输入
	for input.Scan() {
		msg := input.Text()
		if msg == "/quit" {
			break
		}
		msgChannel <- message{
			client,
			msg,
		}
	}

	// 用户退出
	msgChannel <- message{
		client,
		fmt.Sprintf("%s 离开了聊天室", client.username),
	}
	leaving <- client
	conn.Close()
}

func clientWriter(conn net.Conn, client client) {
	for msg := range client.ch {
		fmt.Fprintln(conn, msg)
	}
}

func handleMessage(m message) {
	if m.content == "" {
		return
	}
	if strings.HasPrefix(m.content, "/") {
		// 处理命令
		parts := strings.SplitN(m.content, " ", 2)
		cmd := parts[0]
		switch cmd {
		case "/list":
			var userList []string
			for username := range clients {
				userList = append(userList, username)
			}
			m.c.ch <- fmt.Sprintf("当前在线用户：%s", strings.Join(userList, ", "))
		case "/all":
			if len(parts) > 1 {
				broadcast(fmt.Sprintf("%s: %s", m.c.username, parts[1]), m.c.username)
			}
		default:
			// 处理私聊
			if len(parts) > 1 {
				target := strings.TrimPrefix(cmd, "/")
				if receiver, exists := clients[target]; exists {
					receiver.ch <- fmt.Sprintf("[私聊]%s: %s", m.c.username, parts[1])
					m.c.ch <- fmt.Sprintf("[私聊]发送给 %s: %s", target, parts[1])
				} else {
					m.c.ch <- fmt.Sprintf("用户 %s 不存在", target)
				}
			}
		}
	} else {
		// 默认群发消息
		broadcast(fmt.Sprintf("%s: %s", m.c.username, m.content), m.c.username)
	}
}

func broadcast(msg string, excludeUsername string) {
	for username, client := range clients {
		if username != excludeUsername {
			client.ch <- msg
		}
	}
}
