package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := a[:]
	fmt.Println("反转前的切片：")
	fmt.Println(s)
	s = reverse(s)
	fmt.Println("反转后的切片：")
	fmt.Println(s)
}

func reverse(s []int) []int {
	for x := 0; x < len(s)/2; x++ {
		s[x], s[len(s)-x-1] = s[len(s)-x-1], s[x]
	}
	return s
}
