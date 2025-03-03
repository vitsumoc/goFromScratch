package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// 参数 进行查找的根路径、匹配的文件名
var rootPath string
var keyWord string

// init 函数会最先执行
func init() {
	// 执行参数
	flag.StringVar(&rootPath, "rp", "./", "进行文件查找的根目录（默认当前路径）")
	flag.StringVar(&keyWord, "kw", "", "文件名关键字")
}

func main() {
	flag.Parse()
	// 检查用户输入
	fileInfo, err := os.Stat(rootPath)
	if err != nil || !fileInfo.IsDir() {
		log.Fatalf("输入的路径：%s 不是文件夹。\n", rootPath)
	}
	if len(keyWord) <= 0 {
		log.Fatalf("请输入关键字\n")
	}
	// 递归
	readDir(rootPath, keyWord)
}

func readDir(path string, kw string) {
	// 遍历
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("读取 %s 失败, %v\n", path, err)
	}
	// 处理文件夹中的文件
	for _, file := range files {
		// 是文件夹的
		if file.IsDir() {
			continue
		}
		// 判断关键字
		if strings.Contains(file.Name(), keyWord) {
			fmt.Println(filepath.Join(path, file.Name()))
		}
	}
	// 子文件夹
	for _, file := range files {
		// 非文件夹的
		if !file.IsDir() {
			continue
		}
		// 进入文件夹
		subPath := filepath.Join(path, file.Name())
		readDir(subPath, keyWord)
	}
}
