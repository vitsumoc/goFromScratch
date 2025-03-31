package main

import (
	"fmt"
	"time"
)

func main() {
	var balance int = 1000 // 银行账户有 1000 元
	for x := 0; x < 1000; x++ {
		editBalance(&balance, 50) // 存入 50 元
		time.Sleep(1 * time.Millisecond)
		editBalance(&balance, -50) // 取出 50 元
		time.Sleep(1 * time.Millisecond)
	}
	// 打印余额
	fmt.Printf("存取款 1000 次后，账户余额是 %d 元。\n", balance)
}

func editBalance(balance *int, n int) {
	*balance = *balance + n
}
