package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./printMyself.go")
	if err != nil {
		fmt.Printf("打开文件出错: %v\n", err)
		return
	}
	defer file.Close()

	// 读取所有内容
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("读取文件出错: %v\n", err)
		return
	}
	// 打印内容
	fmt.Printf(string(content))
}
