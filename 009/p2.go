package main

import "fmt"

func main() {
	var a int = 3
	// &变量 => 取地址 => 获得变量的地址
	var p *int = &a
	// *指针 => 解指针 => 获得指针指向的变量
	fmt.Printf("a 的地址：%p，a 的值：%d\n", &a, a)
	fmt.Printf("a 的地址：%p，a 的值：%d\n", p, *p)
}
