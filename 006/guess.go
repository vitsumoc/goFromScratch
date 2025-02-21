package main

import "fmt"

func main() {
	const answer int = 11451
	var guess int
	fmt.Println("您有 100 次机会，请猜测一个 1000000 以内的整数：")
	for attempts := 99; attempts >= 0; attempts = attempts - 1 {
		fmt.Scanln(&guess)
		if guess == answer {
			fmt.Println("恭喜您！猜中了！")
			return
		} else if guess < answer {
			fmt.Println("您输入的数字小了，您的剩余机会：", attempts)
		} else {
			fmt.Println("您输入的数字大了，您的剩余机会：", attempts)
		}
	}
	fmt.Println("很遗憾，您没有猜中，正确答案是：", answer)
}
