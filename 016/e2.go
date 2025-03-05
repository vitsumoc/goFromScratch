package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := myOpen("这个文件肯定不存在")
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	defer file.Close()
}

// 故障上报 集中解决
func myOpen(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}
