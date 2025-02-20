package main

import "fmt"

func main() {
	const length float64 = 5
	const width float64 = 3

	var area float64 = length * width
	var perimeter float64 = 2 * (length + width)

	fmt.Println(area)
	fmt.Println(perimeter)
}
