package main

import "fmt"

func main() {
	name := [4]string{"大雄", "胖虎", "出木杉", ""}
	age := [4]int{8, 8, 8, 0}
	score := [4]int{0, 30, 100, 0}
	fmt.Printf("请输入您的姓名：")
	fmt.Scanln(&name[3])
	fmt.Printf("请输入您的年龄：")
	fmt.Scanln(&age[3])
	fmt.Printf("请输入您的考试成绩：")
	fmt.Scanln(&score[3])
	// 输出
	fmt.Printf("|姓名\t|年龄\t|成绩\t|\n")
	fmt.Printf("|-------|-------|-------|\n")
	for x := 0; x < len(name); x++ {
		fmt.Printf("|%s\t|%d\t|%d\t|\n", name[x], age[x], score[x])
	}
}
