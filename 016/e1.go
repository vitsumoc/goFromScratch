package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("这个文件肯定不存在")
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	defer file.Close()
}
