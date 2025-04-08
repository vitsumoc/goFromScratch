// channel 语法
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 基本语法示例
	var ch1 chan int = make(chan int, 5)
	var ch2 chan int = make(chan int)
	fmt.Println(ch1)
	fmt.Println(ch2)
	fmt.Println(reflect.TypeOf(ch1))
	fmt.Println(reflect.TypeOf(ch2))
	ch1 <- 100
	var a int = <-ch1
	fmt.Println(a)
	close(ch1)
}

func onlyWrite(ch chan<- int) {

}

func onlyRead(ch <-chan int) {

}
