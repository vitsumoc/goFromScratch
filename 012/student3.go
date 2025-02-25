package main

import "fmt"

type sScore struct {
	chinese float64 // 语文成绩
	math    float64 // 数学成绩
	english float64 // 英语成绩
}

type student struct {
	name   string  // 姓名
	age    int     // 年龄
	number string  // 联系电话
	height float64 // 身高
	weight float64 // 体重
	score  sScore  // 考试成绩
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
		weight: 30,
		score: sScore{
			chinese: 90,
			math:    75,
			english: 80,
		},
	}
	// 打印学生成绩
	fmt.Printf("姓名\t语文\t数学\t英语\n")
	var jx = stuMap[9987]
	fmt.Printf("%s\t%.1f\t%.1f\t%.1f\n", jx.name, jx.score.chinese, jx.score.math, jx.score.english)
}
