package main

import "fmt"

func main() {
	// 将 s 初始化为 0 长度的数组
	s := make([]int, 0)
	fmt.Printf("长度：%d，容量：%d\n", len(s), cap(s))
	// 塞 6 个数据进去
	for x := 0; x < 6; x++ {
		s = append(s, x)
	}
	fmt.Printf("长度：%d，容量：%d\n", len(s), cap(s))
	// 在放 10 个数据进去
	for x := 0; x < 10; x++ {
		s = append(s, x)
	}
	fmt.Printf("长度：%d，容量：%d\n", len(s), cap(s))
}
