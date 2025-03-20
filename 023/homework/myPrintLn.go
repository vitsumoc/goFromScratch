package main

import (
	"fmt"
	"reflect"
)

func myPrintLn(a ...any) {
	for _, _a := range a {
		// 支持整数
		if intA, ok := _a.(int); ok {
			fmt.Printf("%d", intA)
			continue
		}
		// 支持字符串
		if strA, ok := _a.(string); ok {
			fmt.Printf("%s", strA)
			continue
		}
		// 支持整数切片
		if intsA, ok := _a.([]int); ok {
			fmt.Printf("len=%d;value=[", len(intsA))
			for _, i := range intsA {
				fmt.Printf("%d ", i)
			}
			fmt.Printf("];")
			continue
		}
		// 其他的数据都不支持，直接输出类型和地址
		fmt.Printf("type=%s;addr=%p;", reflect.TypeOf(_a), &a)
	}
	defer fmt.Printf("\n")
}

func main() {
	myPrintLn(3)              // 支持整数
	myPrintLn("你好世界")         // 支持字符串
	myPrintLn([]int{1, 2, 3}) // 支持整数切片
	// 对于不支持的数据，会输出类型和地址
	myPrintLn([]string{"你好", "世界"}) // 不支持字符串切片
}
