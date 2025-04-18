package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8864") // 监听 8864 端口
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
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close() // 网络链接需要关闭
	defer fmt.Println("客户端退出")
	for {
		buf := make([]byte, 1024)
		n, err := c.Read(buf)
		if err != nil {
			return
		}
		msg := string(buf[:n])
		if err != nil {
			return
		}
		// 回声
		_, err = c.Write([]byte(msg))
		if err != nil {
			return
		}
	}
}
