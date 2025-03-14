// 查看 uint 类型的内存表示
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a uint8 = 114

	fmt.Printf("uint8 类型的 a 的内存长度为：%d byte\n", unsafe.Sizeof(a))
	fmt.Printf("uint8 类型的 a 的内存表示为：%08b\n", a)
}
