package main

import "fmt"

// 定义 平面中的一个点
type point struct {
	x float64
	y float64
}

func main() {
	// 声明
	var p1 point
	// 初始化
	p1 = point{
		x: 3,
		y: 14,
	}
	// 赋值
	p1.x = 5
	// 取值
	fmt.Println(p1.x)
}
