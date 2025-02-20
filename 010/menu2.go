package main

import "fmt"

func main() {
	name := [4]string{"红烧鲫鱼", "毛豆烧鸡", "家常豆腐", "青椒炒蛋"}
	price := [4]int{28, 22, 16, 14}

	fmt.Println("菜单")
	fmt.Printf("序号\t菜品\t价格\t\n")
	for x := 0; x < len(name); x++ {
		fmt.Printf("%d\t%s\t%d元\t\n", x, name[x], price[x])
	}
}
