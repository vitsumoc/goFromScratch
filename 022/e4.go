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

func (c *Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	p1 := Point{3, 4}
	c1 := Circle{p1, 1}
	fmt.Printf("圆 c1 的面积是: %.6f\n", c1.Area())
}
