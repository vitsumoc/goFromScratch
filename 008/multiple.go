package main

import "fmt"

func main() {
	var a, b int = 3, 5
	fmt.Println(a, b)
	// 使用 swap 函数, 交换 a, b 的值
	a, b = swap(a, b)
	fmt.Println(a, b)
}

// 将输入参数交换次序返回
func swap(a int, b int) (int, int) {
	return b, a
}
