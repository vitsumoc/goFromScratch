package main

import "fmt"

func main() {
	var a [8]int
	a[0] = 100
	fmt.Println(a[0])
	var b [7]int = [7]int{9, 9, 9, 8, 8, 7, 7}
	c := [7]int{1, 2, 3, 4, 5, 6, 7}
	d := [...]int{3, 7, 2, 1}
}
