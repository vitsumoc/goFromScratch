package main

import (
	"fmt"
	"time"
)

func main() {
	var counter int
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				counter++
			}
		}()
	}
	time.Sleep(3 * time.Second)
	fmt.Println("最终计数值:", counter)
}
