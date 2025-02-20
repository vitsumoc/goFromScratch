package main

import "fmt"

func main() {
	for a := 0; a < 3; a = a + 1 {
		fmt.Println("当前 a 的值为：", a)
	}
}
