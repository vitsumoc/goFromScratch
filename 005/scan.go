package main

import "fmt"

func main() {
	var i int
	var s string
	fmt.Println("请输入一个整数：")
	fmt.Scanln(&i)
	fmt.Println("请输入一个字符串：")
	fmt.Scanln(&s)
	fmt.Println("您输入的整数是：")
	fmt.Println(i)
	fmt.Println("您输入的字符串是：")
	fmt.Println(s)
}
