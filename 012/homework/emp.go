package main

import "fmt"

type employee struct {
	name   string  // 姓名
	gender string  // 性别
	age    int     // 年龄
	salary float64 // 薪水
}

func main() {
	// 初始化员工数组
	emps := [4]employee{}
	// 填入信息
	emps[0] = employee{name: "张一猩", gender: "男", age: 19, salary: 4900}
	emps[1] = employee{name: "闻大磊", gender: "男", age: 28, salary: 8000}
	emps[2] = employee{name: "刘乐之", gender: "女", age: 24, salary: 8000}
	emps[3] = employee{name: "王叶伦", gender: "男", age: 26, salary: 12000}
	// 员工的平均薪水
	fmt.Printf("员工的平均薪水是：%.2f\n", avgSalary(&emps))
	// 最年长的员工
	oldEmp := oldest(&emps)
	fmt.Printf("最年长的员工是 %s, 年龄是 %d 岁。\n", oldEmp.name, oldEmp.age)
	// 男性员工的比例
	fmt.Printf("男性员工占比：%.1f%%\n", maleProp(&emps))
}

func avgSalary(emps *[4]employee) float64 {
	var sumSalary float64 = 0
	for _, emp := range emps {
		sumSalary += emp.salary
	}
	return sumSalary / float64(len(emps))
}

func oldest(emps *[4]employee) *employee {
	var oldestEmp *employee = &emps[0]
	for x := 0; x < len(emps); x++ {
		if emps[x].age > oldestEmp.age {
			oldestEmp = &emps[x]
		}
	}
	return oldestEmp
}

func maleProp(emps *[4]employee) float64 {
	var maleCount int = 0
	for _, emp := range emps {
		if emp.gender == "男" {
			maleCount += 1
		}
	}
	return float64(maleCount) * 100.0 / float64(len(emps))
}
