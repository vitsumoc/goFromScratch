package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建两个channel
	ch1 := make(chan string)
	ch2 := make(chan string)

	// 启动两个goroutine，分别向不同的channel发送消息
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(1 * time.Second)
			ch1 <- fmt.Sprintf("来自ch1的消息 %d", i)
		}
		close(ch1)
	}()

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(2 * time.Second)
			ch2 <- fmt.Sprintf("来自ch2的消息 %d", i)
		}
		close(ch2)
	}()

	// 使用select多路复用处理两个channel
	for {
		select {
		case msg1, ok := <-ch1:
			if !ok {
				fmt.Println("ch1 已关闭")
				ch1 = nil // 将channel设置为nil，这样select就不会再尝试从这个channel接收
			} else {
				fmt.Println(msg1)
			}
		case msg2, ok := <-ch2:
			if !ok {
				fmt.Println("ch2 已关闭")
				ch2 = nil
			} else {
				fmt.Println(msg2)
			}
		default:
			// 当所有channel都关闭时退出循环
			if ch1 == nil && ch2 == nil {
				fmt.Println("所有channel都已关闭，程序退出")
				return
			}
			// 如果没有消息可接收，短暂休眠
			time.Sleep(100 * time.Millisecond)
		}
	}
}
