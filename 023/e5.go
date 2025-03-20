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

// 多边形
type Polygon struct {
	Points []Point
}

// 多边形没有计算面积的 Area 方法

// 定义 AreaAble 接口，表示所有可以计算面积的类型
type AreaAble interface {
	Area() float64
}

// 计价程序
func calPrice(sharp any) float64 {
	areaSharp, ok := sharp.(AreaAble)
	if ok {
		return areaSharp.Area() * 2
	}
	return 200
}

func main() {
	p1 := Point{3, 4}
	c1 := Circle{p1, 1}
	fmt.Printf("圆 c1 的价格是: %.2f\n", calPrice(&c1))
	r1 := Rect{p1, 5, 5}
	fmt.Printf("矩形 r1 的价格是: %.2f\n", calPrice(&r1))
	polygon1 := Polygon{[]Point{{0, 0}, {0, 3}, {4, 0}}}
	fmt.Printf("多边形 polygon1 的价格是: %.2f\n", calPrice(&polygon1))
}
