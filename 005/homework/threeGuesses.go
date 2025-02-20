package main

import "fmt"

func main() {
	const answer int = 7
	var guess int
	fmt.Println("请猜测一个 10 以内的整数，您有三次机会：")
	fmt.Scanln(&guess)
	if guess == answer {
		fmt.Println("恭喜您！猜中了！")
		return
	} else if guess < answer {
		fmt.Println("您输入的数字小了，您还有两次机会：")
	} else {
		fmt.Println("您输入的数字大了，您还有两次机会：")
	}
	fmt.Scanln(&guess)
	if guess == answer {
		fmt.Println("恭喜您！猜中了！")
		return
	} else if guess < answer {
		fmt.Println("您输入的数字小了，您还有一次机会：")
	} else {
		fmt.Println("您输入的数字大了，您还有一次机会：")
	}
	fmt.Scanln(&guess)
	if guess == answer {
		fmt.Println("恭喜您！猜中了！")
		return
	} else {
		fmt.Println("很遗憾，您没有猜中，正确答案是：", answer)
	}
}
