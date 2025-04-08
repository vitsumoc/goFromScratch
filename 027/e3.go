package main

import (
	"fmt"
	"sync"
	"time"
)

var balanceLocker sync.Mutex

func main() {
	var balance int = 1000 // 银行账户有 1000 元
	for x := 0; x < 500; x++ {
		go func() {
			for y := 0; y < 2; y++ {
				editBalance(&balance, 50) // 存入 50 元
				time.Sleep(1 * time.Millisecond)
				editBalance(&balance, -50) // 取出 50 元
				time.Sleep(1 * time.Millisecond)
			}
		}()
	}
	time.Sleep(10 * time.Second)
	// 打印余额
	fmt.Printf("存取款 1000 次后，账户余额是 %d 元。\n", balance)
}

func editBalance(balance *int, n int) {
	balanceLocker.Lock()
	*balance = *balance + n
	balanceLocker.Unlock()
}
