package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 要读取的文件路径
	filePath := "./测试文件夹/测试文件.txt"

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("打开文件出错: %v\n", err)
		return
	}

	// 读取所有内容
	content, err := io.ReadAll(file)
	if err != nil {
		file.Close() // 关闭文件
		fmt.Printf("读取文件出错: %v\n", err)
		return
	}
	file.Close() // 关闭文件
	// 打印内容
	fmt.Printf(string(content))
}
