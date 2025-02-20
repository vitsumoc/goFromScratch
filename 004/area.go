package main

import (
	"fmt"
)

func main() {
	var r float64 = 3
	const PI = 3.141592
	area := PI * r * r
	fmt.Println(area)
	r = 4
	area = PI * r * r
	fmt.Println(area)
}
