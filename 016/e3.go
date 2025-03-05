package main

import (
	"errors"
	"fmt"
	"os"
)

// 自定义错误：文件为空
var ErrEmptyFile = errors.New("文件内容为空")

// readHomeWork 函数用于打开和读取homework.txt文件
func readHomeWork() (string, error) {
	// 读取文件内容
	content, err := os.ReadFile("./homework.txt")
	if err != nil {
		return "", err
	}

	// 检查文件是否为空
	if len(content) == 0 {
		return "", ErrEmptyFile
	}

	return string(content), nil
}

func main() {
	// 调用readHomeWork函数
	content, err := readHomeWork()
	if err != nil {
		fmt.Println("错误：", err)
		return
	}

	// 打印文件内容
	fmt.Println("文件内容：")
	fmt.Println(content)
}
