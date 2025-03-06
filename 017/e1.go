package main

import (
	"fmt"
)

func main() {
	fmt.Println("main 函数开始")
	defer f1()
	f2()
	fmt.Println("main 函数结束")
}

func f1() {
	fmt.Println("我是 f1")
}

func f2() {
	fmt.Println("我是 f2")
}
