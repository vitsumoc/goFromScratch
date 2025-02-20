package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Sleep 可以让程序睡觉")
	time.Sleep(5 * time.Second)
	fmt.Println("果然！")
}
