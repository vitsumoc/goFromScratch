package main

import "fmt"

func main() {
	var a int = 3
	var pa *int = &a
	var ppa **int = &pa
	fmt.Printf("a：%v\npa：%v\nppa：%v\n", a, pa, ppa)
}
