package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var counter int
	var counterLocker sync.Mutex
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				counterLocker.Lock()
				counter++
				counterLocker.Unlock()
			}
		}()
	}
	time.Sleep(3 * time.Second)
	fmt.Println("最终计数值:", counter)
}
