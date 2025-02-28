package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 绝对路径
	const absolutePath = "C:\\Users\\vc\\goFromScratch\\src\\014\\测试文件夹"
	// 相对路径
	const relativePath = "./测试文件夹"
	// 按路径获得文件夹信息
	fileInfo, err := os.Stat(relativePath)
	if err != nil {
		log.Fatal(err)
	}
	// 打印文件夹信息
	fmt.Printf("fileInfo.Name()：%v\n", fileInfo.Name())       // 文件名称
	fmt.Printf("fileInfo.Size()：%v\n", fileInfo.Size())       // 文件大小
	fmt.Printf("fileInfo.Mode()：%v\n", fileInfo.Mode())       // 读写权限
	fmt.Printf("fileInfo.ModTime()：%v\n", fileInfo.ModTime()) // 修改时间
	fmt.Printf("fileInfo.IsDir()：%v\n", fileInfo.IsDir())     // 是否是文件夹

	if !fileInfo.IsDir() {
		log.Fatal("非文件夹")
	}

	// 遍历文件夹内容
	files, err := os.ReadDir(relativePath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
