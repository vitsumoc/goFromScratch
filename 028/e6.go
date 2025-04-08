package main

import "fmt"

func main() {
	var ch1 chan int = make(chan int, 3)
	fmt.Printf("ch1的长度：%d\n", len(ch1))
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	fmt.Printf("ch1的长度：%d\n", len(ch1))
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
	fmt.Printf("ch1的长度：%d\n", len(ch1))
}
