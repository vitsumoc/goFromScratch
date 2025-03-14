// 查看 int 类型的内存表示
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int8 = -114

	aPtr := (*uint8)(unsafe.Pointer(&a))

	fmt.Printf("int8 类型的 a 的内存长度为：%d byte\n", unsafe.Sizeof(a))
	fmt.Printf("int8 类型的 a 的内存表示为：%08b\n", *aPtr)
}
