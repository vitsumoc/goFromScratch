package main

import "fmt"

func main() {
	name := [4]string{"红烧鲫鱼", "毛豆烧鸡", "家常豆腐", "青椒炒蛋"}
	price := [4]int{28, 22, 16, 14}

	for {
		showMenu(name, price)
		editPrice(&price)
	}
}

// 显示菜单
func showMenu(nameArr [4]string, priceArr [4]int) {
	fmt.Println("菜单")
	fmt.Printf("序号\t菜品\t价格\t\n")
	for x := 0; x < len(nameArr); x++ {
		fmt.Printf("%d\t%s\t%d元\t\n", x, nameArr[x], priceArr[x])
	}
}

// 修改价格函数，让用户输入序号和价格，随后赋值
func editPrice(priceArray *[4]int) {
	var i, p int
	fmt.Printf("请输入需修改价格的菜品序号：")
	fmt.Scanln(&i)
	fmt.Printf("请输入需菜品价格：")
	fmt.Scanln(&p)
	(*priceArray)[i] = p
}
