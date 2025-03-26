package main

import (
	"fmt"
	"io"
	"log"
	"net"
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
	_, err := io.WriteString(c, "焯！这不就来了吗？\n")
	if err != nil {
		return
	}
	for {
		buf := make([]byte, 1024)
		n, err := c.Read(buf)
		if err != nil {
			return
		}
		msg := string(buf[:n])
		if msg == "你好\n" {
			_, err = io.WriteString(c, "咳咳，优雅，太优雅了～\n")
			continue
		}
		if msg == "你好？\n" {
			_, err = io.WriteString(c, "奥力给！干了兄弟们！\n")
			continue
		}
		if msg == "hello？\n" {
			_, err = io.WriteString(c, "awsl！这 hello 开眼了属于是\n")
			continue
		}
		if msg == "你干什么呢？\n" {
			_, err = io.WriteString(c, "你~干~嘛~啊\n")
			continue
		}
		if msg == "再见\n" {
			_, err = io.WriteString(c, "溜了溜了~\n")
			return
		}
		_, err = io.WriteString(c, "啊对对对！\n")
		if err != nil {
			return
		}
	}
}
