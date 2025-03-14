package main

import "fmt"

func main() {
	var a uint8 = 3
	var b uint8 = ^a
	fmt.Printf("b 的值是：%d\n", b)
	fmt.Printf("b 的内存表示为：%08b\n", b)
}
