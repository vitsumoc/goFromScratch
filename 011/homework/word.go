// 统计字符和频率
package main

import "fmt"

func main() {
	var paragraph string
	fmt.Println("请输入一段文本：")
	fmt.Scanln(&paragraph)
	// 字符总量、字符map
	var allCount float64 = 0
	charMap := make(map[rune]int)
	// 遍历字符串
	for _, char := range paragraph {
		// 统计总数
		allCount += 1
		// 统计单个字符数
		_, ok := charMap[char]
		if !ok {
			charMap[char] = 1
		} else {
			charMap[char] += 1
		}
	}
	// 制表输出
	fmt.Printf("|字符\t|次数\t|频率\t|\n")
	fmt.Printf("|-------|-------|-------|\n")
	for char, count := range charMap {
		fmt.Printf("|%c\t|%d\t|%.2f%%\n", char, count, float64(count)*100/allCount)
	}
}
