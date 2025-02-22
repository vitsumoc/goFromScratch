package main

import "fmt"

func main() {
	var n int
	var count int
	fmt.Printf("请输入需要计算的项数：")
	fmt.Scanln(&n)
	fmt.Printf("斐波那契数列第 %d 项的值为：%d，递归计算调用次数为：%d", n, fibonacci(n, &count), count)
}

func fibonacci(n int, count *int) int {
	*count += 1
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	return fibonacci(n-1, count) + fibonacci(n-2, count)
}
