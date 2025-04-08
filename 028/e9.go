package main

import "fmt"

func main() {
	var ch1 chan int = make(chan int)
	close(ch1)
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)

	var ch2 chan int = make(chan int)
	close(ch2)
	i, ok := <-ch2
	fmt.Println(i, ok)

	var ch3 chan int = make(chan int)
	close(ch3)
	ch3 <- 1
	fmt.Println(<-ch3)
}
