package main

import (
	"fmt"
	"time"
)

func main() {
	var ch1 chan int = make(chan int)
	go func() {
		for {
			ch1 <- 1
			time.Sleep(time.Second * 1)
		}
	}()

	for {
		fmt.Println(<-ch1)
	}
}
