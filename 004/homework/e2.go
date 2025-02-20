package main

import "fmt"

func main() {
	const APPLE_PRICE float64 = 5.5
	const BANANA_PRICE float64 = 2.95

	var appleWeight float64 = 5
	var bananaWeight float64 = 3

	var sum float64 = appleWeight*APPLE_PRICE + bananaWeight*BANANA_PRICE
	fmt.Println(sum)
}
