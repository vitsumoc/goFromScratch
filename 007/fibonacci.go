package main

import "fmt"

func main() {
	fmt.Println("斐波那契数列第 10 项目的值为：", fibonacci(10))
}

func fibonacci(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
