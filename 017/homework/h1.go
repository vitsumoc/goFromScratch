package main

import "fmt"

func main() {
	defer fmt.Print("A")
	defer fmt.Print("B")
	fmt.Print("C")
	return
	defer fmt.Print("D")
}
