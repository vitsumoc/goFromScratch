package main

import "fmt"

func main() {
	// 创建数组 a，切片s1，切片s2
	a := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	s1 := a[:]
	s2 := s1[3:5]
	// 修改 s2
	s2[0] = 9
	// 发现都被修改了
	fmt.Println(a)
	fmt.Println(s1)
	fmt.Println(s2)
}
