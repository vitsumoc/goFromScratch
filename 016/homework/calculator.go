package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	fmt.Println("欢迎使用计算器！")
	fmt.Println("请输入计算表达式（格式：数字 运算符 数字），支持的运算符有 + - * / %")
	fmt.Println("输入 'q' 退出程序")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "q" {
			fmt.Println("感谢使用，再见！")
			break
		}

		if input == "" {
			continue
		}

		result, err := calculate(input)
		if err != nil {
			fmt.Printf("错误: %v\n", err)
		} else {
			fmt.Printf("结果: %v\n", result)
		}
	}
}

func calculate(input string) (float64, error) {
	// 分割输入字符串
	parts := strings.Fields(input)
	if len(parts) != 3 {
		return 0, fmt.Errorf("输入格式错误，请使用格式：数字 运算符 数字")
	}

	// 解析第一个数字
	num1, err := parseNumber(parts[0])
	if err != nil {
		return 0, fmt.Errorf("第一个数字无效: %v", err)
	}

	// 获取运算符
	operator := parts[1]

	// 解析第二个数字
	num2, err := parseNumber(parts[2])
	if err != nil {
		return 0, fmt.Errorf("第二个数字无效: %v", err)
	}

	// 执行计算
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("除数不能为零")
		}
		return num1 / num2, nil
	case "%":
		// 检查是否为整数
		if math.Floor(num1) != num1 || math.Floor(num2) != num2 {
			return 0, fmt.Errorf("取模运算只支持整数")
		}
		if num2 == 0 {
			return 0, fmt.Errorf("取模运算的除数不能为零")
		}
		return float64(int(num1) % int(num2)), nil
	default:
		return 0, fmt.Errorf("不支持的运算符: %s", operator)
	}
}

func parseNumber(s string) (float64, error) {
	var num float64
	_, err := fmt.Sscanf(s, "%f", &num)
	if err != nil {
		return 0, fmt.Errorf("无法解析数字")
	}
	return num, nil
}
