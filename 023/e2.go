package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

type Circle struct {
	p Point   // 圆心点
	r float64 // 半径
}

// 圆面积
func (c *Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

type Rect struct {
	BottomLeft Point   // 左下顶点
	Width      float64 // 宽度
	Height     float64 // 长度
}

// 矩形面积
func (r *Rect) Area() float64 {
	return r.Width * r.Height
}

// 计价程序 (传入圆形，计算价格)
func calPrice(c Circle) float64 {
	return c.Area() * 2
}

func main() {
	p1 := Point{3, 4}
	c1 := Circle{p1, 1}
	fmt.Printf("圆 c1 的价格是: %.2f\n", calPrice(c1))
	r1 := Rect{p1, 5, 5}
	fmt.Printf("矩形 r1 的价格是: %.2f\n", calPrice(r1))
}
