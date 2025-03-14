package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a bool = true
	var b int = 1
	var c float64 = 1.0

	fmt.Printf("布尔类型 a 的内存长度为：%d byte\n", unsafe.Sizeof(a))
	fmt.Printf("整数类型 b 的内存长度为：%d byte\n", unsafe.Sizeof(b))
	fmt.Printf("浮点数类型 c 的内存长度为：%d byte\n", unsafe.Sizeof(c))
}
