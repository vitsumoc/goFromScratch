package main

import (
	"fmt"
	"os"
)

func main() {
	// 声明和初始化
	var contacts map[string]string
	contacts = make(map[string]string)
	// 赋值
	contacts["大雄"] = "230-3333"
	contacts["胖虎"] = "230-3065"
	contacts["静香"] = "230-9987"
	contacts["小夫"] = "230-3551"
	contacts["出木杉"] = "230-7775"
	// 选中的功能菜单
	var menu int = 0
	// 程序主循环
	for {
		fmt.Println("请选择功能：")
		fmt.Println("1. 查看通讯录")
		fmt.Println("2. 查找联系人")
		fmt.Println("3. 新增联系人")
		fmt.Println("4. 删除联系人")
		fmt.Println("5. 修改联系人")
		fmt.Println("0. 退出程序")
		fmt.Scanln(&menu)
		checkMenu(menu, contacts)
		// 用户敲击回车进入下一个菜单
		fmt.Printf("按回车继续...")
		fmt.Scanln()
		menu = 0
	}
}

// 选择功能
func checkMenu(m int, c map[string]string) {
	// 退出程序
	if m == 0 {
		os.Exit(0)
	}
	// 列表查看
	if m == 1 {
		list(c)
		return
	}
	// 检索单个
	if m == 2 {
		var name string
		fmt.Println("请输入联系人姓名：")
		fmt.Scanln(&name)
		query(c, name)
		return
	}
	// 新增
	if m == 3 {
		var name string
		var number string
		fmt.Println("请输入联系人姓名：")
		fmt.Scanln(&name)
		fmt.Println("请输入电话：")
		fmt.Scanln(&number)
		add(c, name, number)
		return
	}
	if m == 4 {
		var name string
		fmt.Println("请输入联系人姓名：")
		fmt.Scanln(&name)
		del(c, name)
		return
	}
	if m == 5 {
		var name string
		var number string
		fmt.Println("请输入联系人姓名：")
		fmt.Scanln(&name)
		fmt.Println("请输入电话：")
		fmt.Scanln(&number)
		edit(c, name, number)
		return
	}
	fmt.Println("您输入的功能代码有误。")
}

// 查看通讯录
func list(c map[string]string) {
	fmt.Printf("|姓名\t|电话\t|\n")
	fmt.Printf("|-------|-------|\n")
	for name, number := range c {
		fmt.Printf("|%s\t|%s\t|\n", name, number)
	}
}

// 查找联系人
func query(c map[string]string, name string) {
	number, ok := c[name]
	// 联系人不存在
	if !ok {
		fmt.Printf("联系人 %s 不存在。\n", name)
		return
	}
	// 联系人存在，正确输出
	fmt.Printf("|姓名\t|电话\t|\n")
	fmt.Printf("|-------|-------|\n")
	fmt.Printf("|%s\t|%s\t|\n", name, number)
}

// 添加联系人
func add(c map[string]string, name string, number string) {
	_, ok := c[name]
	if ok {
		fmt.Printf("联系人 %s 已经存在。\n", name)
		return
	}
	c[name] = number
	fmt.Printf("联系人 %s 添加成功。\n", name)
}

// 删除联系人
func del(c map[string]string, name string) {
	_, ok := c[name]
	// 联系人不存在
	if !ok {
		fmt.Printf("联系人 %s 不存在。\n", name)
		return
	}
	delete(c, name)
	fmt.Printf("联系人 %s 已删除。\n", name)
}

// 修改联系人
func edit(c map[string]string, name string, number string) {
	_, ok := c[name]
	if !ok {
		fmt.Printf("联系人 %s 不存在。\n", name)
		return
	}
	c[name] = number
	fmt.Printf("联系人 %s 修改成功。\n", name)
}
