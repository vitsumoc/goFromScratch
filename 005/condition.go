package main

import "fmt"

func main() {
	if 1+1 == 2 {
		fmt.Println("1+1==2 为真：此句执行")
	}

	if false {
		fmt.Println("false 为假：此句不执行")
	} else {
		fmt.Println("上一条判断语句为假：此句执行")
	}

	if 1+1 != 2 {
		fmt.Println("1+1!=2 为假：此句不执行")
	} else if 3.14 > 3.1 {
		fmt.Println("3.14 > 3.1 为真：此句执行")
	} else {
		fmt.Println("上一条判断语句为真：此句不执行")
	}
}
