package main

import "fmt"

const PI float64 = 3.14159

func main() {
	var r float64
	for {
		fmt.Println("请输入球的半径，输入 0 退出：")
		fmt.Scanln(&r)
		if r == 0 {
			return
		}
		fmt.Printf("半径为 %.2f 的球体，其体积为：%.2f\n", r, volume(r))
	}
}

func volume(r float64) float64 {
	return PI * r * r * r * 4 / 3
}
