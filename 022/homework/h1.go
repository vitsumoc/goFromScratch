package main

import (
	"fmt"
)

type Point struct {
	X float64
	Y float64
}

type Rect struct {
	BottomLeft Point
	Width      float64
	Height     float64
}

func (r *Rect) Area() float64 {
	return r.Width * r.Height
}

func main() {
	p1 := Point{3, 4}
	r1 := Rect{p1, 5, 5}
	fmt.Printf("矩形 r1 的面积是: %.6f\n", r1.Area())
}
