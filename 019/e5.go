// 查看 float 类型的内存表示
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a float32 = 3.1415925
	aPtr := (*uint32)(unsafe.Pointer(&a))
	fmt.Printf("float32 类型的 a 的内存长度为：%d byte\n", unsafe.Sizeof(a))
	fmt.Printf("float32 类型的 a 的内存表示为：%032b\n", *aPtr)
}
