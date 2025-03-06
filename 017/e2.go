package main

import "fmt"

func main() {
	fmt.Println("main 函数开始")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	defer fmt.Println(4)
	defer fmt.Println(5)
	fmt.Println("main 函数结束")
}
