package main

import "fmt"

func main() {
	if true {
		// 代码块
		s1 := "hello world"
		fmt.Println(s1)
	}
	// fmt.Println(s1)

	f1()
	// fmt.Println(s2)

	fmt.Println(s3)
}

// 函数
func f1() {
	s2 := "我是s2"
	fmt.Println(s2)
}

// 包
var s3 string = "我是s3"

// 可导出
var S4 string = "我是s4"
