package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8864")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("服务启动，监听：0.0.0.0:8864")
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

var addr2Conn = make(map[string]net.Conn)

func handleConn(c net.Conn) {
	var who string = c.RemoteAddr().String()
	defer c.Close() // 网络链接需要关闭
	defer fmt.Println("客户端退出")
	defer delete(addr2Conn, who)
	addr2Conn[who] = c
	input := bufio.NewScanner(c)
	for input.Scan() {
		msg := who + ":" + input.Text() + "\n"
		for _, conn := range addr2Conn {
			_, err := conn.Write([]byte(msg))
			if err != nil {
				log.Println("发送消息失败:", err)
				continue
			}
		}
	}
}
