// 查看 bool 类型的内存表示
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a bool = true
	var b bool = false

	aPtr := (*uint8)(unsafe.Pointer(&a))
	bPtr := (*uint8)(unsafe.Pointer(&b))

	fmt.Printf("布尔类型 a 的内存长度为：%d byte\n", unsafe.Sizeof(a))
	fmt.Printf("布尔类型 a 的内存表示为 %08b\n", *aPtr)
	fmt.Printf("布尔类型 b 的内存表示为 %08b\n", *bPtr)
}
