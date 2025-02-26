package main

import "fmt"

func main() {
	s := "   helloworld!  "
	fmt.Println("strip前的字符串：")
	fmt.Println(s)
	s = strip(s)
	fmt.Println("strip后的字符串：")
	fmt.Println(s)
}

func strip(s string) string {
	left := 0
	right := len(s) - 1
	// 从左往右找到第一个非空格字符的索引
	for left <= right && s[left] == ' ' {
		left++
	}
	// 从右往左找到第一个非空格字符的索引
	for right >= left && s[right] == ' ' {
		right--
	}
	// 返回去除两边空格后的子字符串
	return s[left : right+1]
}
