package main

import "fmt"

func main() {
	// 创建数组和切片
	a := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	s := a[3:5]
	// 修改切片
	s[0] = 9
	// 打印
	fmt.Println(a)
	fmt.Println(s)
	// 查看长度和容量
	fmt.Println(len(s))
	fmt.Println(cap(s))
}
