package main

import (
	"fmt"
	"os"
)

func main() {
	// 文件路径
	const filePath = "./测试文件夹/测试文件.txt"
	// 以写入模式打开文件，如果文件不存在则创建，如果存在则清空内容
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("打开文件出错: %v\n", err)
		return
	}
	// 要写入的内容
	content := "Hello, World!\n这是一个写入文件的示例。"
	// 写入内容到文件
	_, err = file.WriteString(content)
	if err != nil {
		file.Close() // 关闭文件
		fmt.Printf("写入文件出错: %v\n", err)
		return
	}
	file.Close() // 关闭文件
	fmt.Println("文件写入成功！")
}
