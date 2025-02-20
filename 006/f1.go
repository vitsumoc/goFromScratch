package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("这个语句会无限执行，退出请按 Ctrl + C ...")
		time.Sleep(3 * time.Second)
	}
}
