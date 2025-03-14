package main

import "fmt"

func main() {
	var a uint8 = 3
	var b uint8 = 16
	if isOdd(a) {
		fmt.Printf("%d 是奇数\n", a)
	} else {
		fmt.Printf("%d 是偶数\n", a)
	}
	if isOdd(b) {
		fmt.Printf("%d 是奇数\n", b)
	} else {
		fmt.Printf("%d 是偶数\n", b)
	}
}

func isOdd(u uint8) bool {
	return u&1 == 1
}
