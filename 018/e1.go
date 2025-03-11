package main

import "fmt"

func main() {
	var a int = 75
	var b int = 0b1001011
	var c int = 0x4b
	fmt.Printf("a的字面量是 75，值是：%d\n", a)
	fmt.Printf("b的字面量是 0b1001011，值是：%d\n", b)
	fmt.Printf("c的字面量是 0x4b，值是：%d\n", c)
}
