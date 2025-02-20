package main

import "fmt"

func main() {
	// 通用动词示例
	num := 42
	str := "hello"
	fmt.Printf("通用格式输出: %v\n", num)
	fmt.Printf("类型输出: %T\n", str)

	// 整数动词示例
	fmt.Printf("十进制输出: %d\n", num)
	fmt.Printf("二进制输出: %b\n", num)
	fmt.Printf("十六进制小写输出: %x\n", num)
	fmt.Printf("十六进制大写输出: %X\n", num)

	// 浮点数动词示例
	floatNum := 3.14159
	fmt.Printf("默认浮点数输出: %f\n", floatNum)
	fmt.Printf("保留 2 位小数输出: %.2f\n", floatNum)
	fmt.Printf("科学计数法小写输出: %e\n", floatNum)
	fmt.Printf("科学计数法大写输出: %E\n", floatNum)

	// 字符串和字节切片动词示例
	fmt.Printf("字符串输出: %s\n", str)
	fmt.Printf("带引号字符串输出: %q\n", str)

	// 指针动词示例
	ptr := &num
	fmt.Printf("指针地址输出: %p\n", ptr)
}
