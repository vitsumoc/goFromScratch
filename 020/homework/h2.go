package main

import "fmt"

func main() {
	var a uint8 = 233
	fmt.Printf("%d 的二进制表示中有 %d 个1\n", a, count1(a))
}

func count1(i uint8) int {
	var count int = 0
	for x := 0; x < 8; x++ {
		if i&1 == 1 {
			count += 1
		}
		i >>= 1
	}
	return count
}
