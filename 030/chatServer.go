package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	conn     *websocket.Conn
	username string
	send     chan []byte
}

type Message struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Content  string `json:"content"`
	Target   string `json:"target,omitempty"`
}

var (
	clients    = make(map[*Client]bool)
	broadcast  = make(chan Message)
	register   = make(chan *Client)
	unregister = make(chan *Client)
	mu         sync.Mutex
)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket 升级失败:", err)
		return
	}
	log.Println("新客户端连接")

	client := &Client{
		conn: conn,
		send: make(chan []byte, 256),
	}

	register <- client
	log.Println("客户端已注册")

	go handleMessages(client)
	go handleClientSend(client)
}

func handleMessages(client *Client) {
	defer func() {
		log.Println("客户端断开连接:", client.username)
		if client.username != "" {
			// 发送用户退出消息
			broadcast <- Message{
				Type:    "system",
				Content: client.username + " 离开了聊天室",
			}
			// 更新用户列表
			sendUserListToAll()
		}
		unregister <- client
		client.conn.Close()
	}()

	for {
		var msg Message
		err := client.conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("读取消息错误: %v", err)
			break
		}
		log.Printf("收到消息: %+v", msg)

		switch msg.Type {
		case "set_username":
			log.Printf("处理 set_username 消息，当前用户名: %s", msg.Username)
			if isUsernameTaken(msg.Username) {
				// 发送用户名重复消息
				jsonMsg, err := json.Marshal(Message{
					Type:    "error",
					Content: "用户名已存在，请重新设置",
				})
				if err != nil {
					log.Printf("序列化错误消息错误: %v", err)
					continue
				}
				client.send <- jsonMsg
				continue
			}
			client.username = msg.Username
			// 发送系统消息
			systemMsg := Message{
				Type:    "system",
				Content: msg.Username + " 加入了聊天室",
			}
			log.Printf("发送系统消息: %+v", systemMsg)
			broadcast <- systemMsg
			// 立即发送用户列表给所有客户端
			log.Println("准备发送用户列表")
			sendUserListToAll()
		case "message":
			if msg.Target != "" {
				log.Printf("私聊消息: %s -> %s", client.username, msg.Target)
				// 私聊消息
				mu.Lock()
				for c := range clients {
					if c.username == msg.Target || c.username == client.username {
						// 发送给目标用户和发送者
						privateMsg := Message{
							Type:     "private",
							Username: client.username,
							Content:  msg.Content,
							Target:   msg.Target,
						}
						jsonMsg, err := json.Marshal(privateMsg)
						if err != nil {
							log.Printf("序列化私聊消息错误: %v", err)
							continue
						}
						log.Printf("发送私聊消息: %s", string(jsonMsg))
						c.send <- jsonMsg
					}
				}
				mu.Unlock()
			} else {
				log.Printf("广播消息: %s: %s", client.username, msg.Content)
				// 广播消息
				broadcast <- Message{
					Type:     "message",
					Username: client.username,
					Content:  msg.Content,
				}
			}
		default:
			log.Printf("未知消息类型: %s", msg.Type)
		}
	}
}

func handleClientSend(client *Client) {
	defer func() {
		log.Printf("handleClientSend 退出: %s", client.username)
		close(client.send) // 确保 channel 被关闭
		client.conn.Close()
	}()

	for message := range client.send {
		log.Printf("准备发送消息给 %s: %s", client.username, string(message))
		w, err := client.conn.NextWriter(websocket.TextMessage)
		if err != nil {
			log.Printf("NextWriter 错误: %v", err)
			return
		}
		if _, err := w.Write(message); err != nil {
			log.Printf("写入消息错误: %v", err)
			return
		}
		if err := w.Close(); err != nil {
			log.Printf("关闭写入器错误: %v", err)
			return
		}
	}
}

func sendUserListToAll() {
	mu.Lock()
	userList := make([]string, 0, len(clients))
	for client := range clients {
		if client.username != "" {
			userList = append(userList, client.username)
		}
	}
	mu.Unlock()

	log.Printf("当前在线用户: %v", userList)

	// 创建用户列表消息
	userListMsg := Message{
		Type:    "user_list",
		Content: strings.Join(userList, ","),
	}
	jsonMsg, err := json.Marshal(userListMsg)
	if err != nil {
		log.Printf("序列化用户列表错误: %v", err)
		return
	}
	log.Printf("发送用户列表消息: %s", string(jsonMsg))

	// 直接发送给所有客户端
	mu.Lock()
	for client := range clients {
		select {
		case client.send <- jsonMsg:
			log.Printf("用户列表已发送给: %s", client.username)
		default:
			log.Printf("发送用户列表失败，关闭连接: %s", client.username)
			delete(clients, client)
		}
	}
	mu.Unlock()
}

func handleBroadcast() {
	for {
		msg := <-broadcast
		jsonMsg, err := json.Marshal(msg)
		if err != nil {
			log.Printf("error marshaling message: %v", err)
			continue
		}

		mu.Lock()
		for client := range clients {
			select {
			case client.send <- jsonMsg:
			default:
				close(client.send)
				delete(clients, client)
			}
		}
		mu.Unlock()
	}
}

func handleClientRegistration() {
	for {
		select {
		case client := <-register:
			log.Printf("注册新客户端: %p", client)
			mu.Lock()
			clients[client] = true
			mu.Unlock()
		case client := <-unregister:
			log.Printf("注销客户端: %p, 用户名: %s", client, client.username)
			mu.Lock()
			if _, ok := clients[client]; ok {
				// 先删除用户
				delete(clients, client)
				mu.Unlock()
				// 再发送更新后的用户列表
				sendUserListToAll()
			} else {
				mu.Unlock()
			}
		}
	}
}

func isUsernameTaken(username string) bool {
	mu.Lock()
	defer mu.Unlock()
	for client := range clients {
		if client.username == username {
			return true
		}
	}
	return false
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/ws", handleConnections)

	go handleBroadcast()
	go handleClientRegistration()

	log.Println("服务器启动在 :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
