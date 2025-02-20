package main

import "fmt"

func main() {
	var a int = 0
	for a < 3 {
		fmt.Println("当前 a 的值为：", a)
		a = a + 1
	}
}
