package main

import "fmt"

func main() {
	fmt.Println(add(1, 1))
	fmt.Println(add(1234567, 9998877))
}

func add(a int, b int) int {
	sum := a + b
	return sum
}
