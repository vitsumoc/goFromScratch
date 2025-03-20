package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (p *Point) Distance() float64 {
	return math.Sqrt(math.Pow(p.x, 2) + math.Pow(p.y, 2))
}

func main() {
	p1 := Point{3, 4}
	fmt.Printf("点 3, 4 到原点的距离是 %.2f\n", p1.Distance())
}
