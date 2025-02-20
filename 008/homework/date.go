package main

import "fmt"

func main() {
	var year int
	var month int
	var day int
	fmt.Println("请依序输入年、月、日：")
	fmt.Scanln(&year)
	fmt.Scanln(&month)
	fmt.Scanln(&day)
	fmt.Printf("您输入日期是：[%d]年[%d]月[%d]日。", year, month, day)
}
