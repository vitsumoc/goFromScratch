package main

import (
	"fmt"
	"log"
	"net"
	"time"
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
		go handleConn(conn)
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
		// 按照回声的方式发送
		if msg != "" {
			// 去掉消息末尾的换行符
			msg = msg[:len(msg)-1]
			go func(msg string) {
				for x := 0; x < 3; x++ {
					// 这里计算小数点的数量，显得更真实一些
					dots := ""
					for i := 0; i <= x; i++ {
						dots += "."
					}
					// 这里把小数点添加到文本之后
					_, err := c.Write([]byte(msg + dots + "\n"))
					if err != nil {
						return
					}
					time.Sleep(1 * time.Second)
				}
			}(msg)
		}
	}
}
