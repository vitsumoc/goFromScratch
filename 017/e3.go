package main

import "os"

func main() {
	handleFile("假设打开什么")
	handleFile2("假设打开什么")
}

func handleFile(path string) {
	file, _ := os.Open(path)
	if false { // 假设某种内容错误
		file.Close()
		return
	}
	if false { // 假设某种内容错误
		file.Close()
		return
	}
	if false { // 假设某种内容错误
		file.Close()
		return
	}
	// 假设全部正确
	file.Close()
	return
}

func handleFile2(path string) {
	file, _ := os.Open(path)
	defer file.Close()
	if false { // 假设某种内容错误
		return
	}
	if false { // 假设某种内容错误
		return
	}
	if false { // 假设某种内容错误
		return
	}
	return
}
