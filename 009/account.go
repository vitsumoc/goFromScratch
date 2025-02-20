package main

import "fmt"

func main() {
	var balance float64 = 1000
	// 花费 500 元
	shopping(balance, 500)
	fmt.Printf("消费 500 元后，静香的账户余额为：%.0f 元", balance)
}

func shopping(balance float64, spent float64) {
	balance = balance - spent
}
