package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (p Point) D1() float64 {
	return math.Sqrt(math.Pow(p.x, 2) + math.Pow(p.y, 2))
}

func (p *Point) D2() float64 {
	return math.Sqrt(math.Pow(p.x, 2) + math.Pow(p.y, 2))
}

func main() {
	p1 := Point{3, 4}
	pp1 := &p1
	fmt.Printf("点 3, 4 到原点的距离是 %.2f\n", p1.D1())
	fmt.Printf("点 3, 4 到原点的距离是 %.2f\n", p1.D2())
	fmt.Printf("点 3, 4 到原点的距离是 %.2f\n", pp1.D1())
	fmt.Printf("点 3, 4 到原点的距离是 %.2f\n", pp1.D2())
}
