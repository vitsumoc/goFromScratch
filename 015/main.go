/*
寻找大文件
默认运行：当前文件夹为根，寻找超过 20M 的文件
带参数运行：
-rp 查找根路径
-s 设置查找的文件大小，单位MB
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// 定义 1mb = 1024kb = 1024 * 1024 byte
const mb int64 = 1 * 1024 * 1024

// 用于记录大文件的结构
type bigFile struct {
	path string // 完整路径
	size int64  // 文件大小
}

// 参数 进行查找的根路径、匹配的文件大小
var rootPath string
var maxSize int64

// 忽略不处理的文件夹
var ignoreDir map[string]bool

// 记录大文件的列表
var bigFileList []bigFile

// init 函数会最先执行
func init() {
	// 执行参数
	flag.StringVar(&rootPath, "rp", "./", "进行文件查找的根目录（默认当前路径）")
	flag.Int64Var(&maxSize, "ms", 20, "查找文件大小（单位MB，默认 20MB）")
	// 跳过不检索的文件夹名
	ignoreDir = make(map[string]bool)
	ignoreDir["node_modules"] = true
	ignoreDir[".svn"] = true
	ignoreDir[".git"] = true
	ignoreDir["__pycache__"] = true
	ignoreDir["pip"] = true
	// 大文件列表
	bigFileList = make([]bigFile, 0)
}

func main() {
	flag.Parse()
	maxSize *= mb
	// 检查用户输入
	fileInfo, err := os.Stat(rootPath)
	if err != nil || !fileInfo.IsDir() {
		log.Fatalf("输入的路径：%s 不是文件夹。\n", rootPath)
	}
	if maxSize < mb {
		log.Fatalf("查找文件大小不能小于 1MB。\n")
	}
	// 在一个递归方法中遍历文件夹
	readDir(rootPath, maxSize)
	// 结果列表输出到文件
	file, err := os.OpenFile(filepath.Join(rootPath, "查找大文件.txt"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer file.Close()
	file.WriteString(fmt.Sprintf("文件路径\t\t\t\t\t\t文件大小\n"))
	for x := 0; x < len(bigFileList); x++ {
		bfInfo := bigFileList[x]
		_, err = file.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t%d\n", bfInfo.path, bfInfo.size))
		if err != nil {
			log.Fatalf("写入结果文件失败, %v\n", err)
		}
	}

	fmt.Println("搜索完成，请查看结果文件！")
}

// 递归调用，读取文件夹和子文件夹，记录超过规定尺寸的文件s
func readDir(path string, maxSize int64) {
	// 遍历
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("读取 %s 失败, %v\n", path, err)
	}
	// 处理文件夹中的文件
	for _, file := range files {
		// 跳过的
		if _, ok := ignoreDir[file.Name()]; ok {
			continue
		}
		// 是文件夹的
		if file.IsDir() {
			continue
		}
		// 判断大小
		info, err := file.Info()
		if err != nil {
			log.Fatalf("读取 %s 中的 %s 失败, %v\n", path, file.Name(), err)
		}
		if info.Size() <= maxSize {
			continue
		}
		// 记录过大的文件
		bigFileList = append(bigFileList, bigFile{
			path: filepath.Join(path, file.Name()),
			size: info.Size(),
		})
	}
	// 子文件夹
	for _, file := range files {
		// 跳过的
		if _, ok := ignoreDir[file.Name()]; ok {
			continue
		}
		// 非文件夹的
		if !file.IsDir() {
			continue
		}
		// 进入文件夹
		subPath := filepath.Join(path, file.Name())
		readDir(subPath, maxSize)
	}
}
