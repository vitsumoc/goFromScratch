package main

import (
	"fmt"
	"time"
)

func main() {
	defer runTime()()
	time.Sleep(time.Second * 5)
}

func runTime() func() {
	start := time.Now()
	fmt.Printf("开始时间：%v\n", start)
	return func() {
		fmt.Printf("结束时间：%v\n", time.Now())
		fmt.Printf("运行时间：%v\n", time.Since(start))
	}
}
