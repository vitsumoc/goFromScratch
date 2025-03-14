package main

import "fmt"

func main() {
	var a uint8 = 3
	var b uint8 = 5
	var c uint8 = a &^ b
	fmt.Printf("c 的值是：%d\n", c)
	fmt.Printf("c 的内存表示为：%08b\n", c)
}
