package main

import "fmt"

func main() {
	var a int = 0
	for {
		fmt.Println("当前 a 的值为：", a)
		a = a + 1
		if a >= 3 {
			break
		}
	}
}
