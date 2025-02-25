package main

import (
	"fmt"
	"math"
)

type point struct {
	x float64
	y float64
}

func main() {
	var p1 point
	var p2 point
	for {
		fmt.Println("请输入第一点坐标：x y")
		fmt.Scanln(&p1.x, &p1.y)
		fmt.Println("请输入第二点坐标：x y")
		fmt.Scanln(&p2.x, &p2.y)
		fmt.Printf("点 %v 和 点 %v 的距离为：%.2f\n\n", p1, p2, distance(p1, p2))
	}
}

func distance(p1 point, p2 point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}
