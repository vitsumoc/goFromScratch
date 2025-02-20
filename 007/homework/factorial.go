package main

import "fmt"

func main() {
	var n int
	fmt.Println("请输入需要计算的值：")
	fmt.Scanln(&n)
	fmt.Printf("%d 的阶乘是：%d! = %d\n", n, n, factorial(n))
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
