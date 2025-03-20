package main

import (
	"fmt"
)

func main() {
	var anyType any
	anyType = "Hello"
	fmt.Println(anyType)
	anyType = 123123
	fmt.Println(anyType)
	anyType = true
	fmt.Println(anyType)
}
