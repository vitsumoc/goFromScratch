package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

func main() {
	fmt.Printf("我是主协程，我的id是：%s\n", goid())

	go func() {
		fmt.Printf("我是子协程，我的id是：%s\n", goid())
		for x := 0; x < 5; x++ {
			fmt.Println(x)
			time.Sleep(1 * time.Second)
		}
		fmt.Println("子协程退出")
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("主协程退出")
}

func goid() string {
	// 获取当前协程ID
	buf := make([]byte, 64)
	n := runtime.Stack(buf, false)
	idField := strings.Fields(string(buf[:n]))[1]
	return idField
}
