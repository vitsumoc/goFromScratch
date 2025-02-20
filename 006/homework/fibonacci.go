// 计算斐波那契数列前一百项的：和、奇数项和、3的倍数项和
package main

import "fmt"

func main() {
	var a int
	var b int
	var sum = 0
	var oddSum = 0
	var tripleSum = 0
	for x := 1; x <= 30; x++ {
		if x == 1 {
			a = 0
			b = 1
		}
		if x > 1 {
			a, b = b, a+b
		}
		sum += a
		if x%2 == 1 {
			oddSum += a
		}
		if x%3 == 0 {
			tripleSum += a
		}
	}
	fmt.Println("所有项的和：", sum)
	fmt.Println("奇数项的和：", oddSum)
	fmt.Println("三倍项的和：", tripleSum)
}
