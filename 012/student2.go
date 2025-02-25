package main

import "fmt"

type student struct {
	name    string  // 姓名
	age     int     // 年龄
	number  string  // 联系电话
	height  float64 // 身高
	weight  float64 // 体重
	chinese float64 // 语文成绩
	math    float64 // 数学成绩
	english float64 // 英语成绩
}

func main() {
	// 声明初始化map
	var stuMap map[int]student
	stuMap = make(map[int]student)

	// 添加学生
	stuMap[9987] = student{
		name:   "静香",
		age:    8,
		number: "230-9987",
		height: 143,
		weight: 39,
	}
	// 打印学生表
	fmt.Printf("姓名\t年龄\t电话\t\t身高\t体重\n")
	var jx student = stuMap[9987]
	fmt.Printf("%s\t%d\t%s\t%.1f\t%.1f\n", jx.name, jx.age, jx.number, jx.height, jx.weight)
}
