package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:7749") // 监听 7749 端口
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("服务启动，监听：0.0.0.0:7749")
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
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
