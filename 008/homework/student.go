package main

import "fmt"

func main() {
	var name string
	var age int
	var score int
	fmt.Printf("请输入您的姓名：")
	fmt.Scanln(&name)
	fmt.Printf("请输入您的年龄：")
	fmt.Scanln(&age)
	fmt.Printf("请输入您的考试成绩：")
	fmt.Scanln(&score)
	// 输出
	fmt.Printf("|姓名\t|年龄\t|成绩\t|\n")
	fmt.Printf("|-------|-------|-------|\n")
	fmt.Printf("|大雄\t|8\t|0\t|\n")
	fmt.Printf("|胖虎\t|8\t|30\t|\n")
	fmt.Printf("|出木杉\t|8\t|100\t|\n")
	fmt.Printf("|%s\t|%d\t|%d\t|\n", name, age, score)
}
