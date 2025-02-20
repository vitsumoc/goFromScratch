package main

import "fmt"

func main() {
	arr := [6]int{5, 3, 8, 1, 2, 4}
	fmt.Printf("排序前的数据：%v\n", arr)
	bubbleSort(&arr)
	fmt.Printf("排序后的数据：%v\n", arr)
}

// bubbleSort 函数实现了冒泡排序算法，用于对整数数组进行排序
func bubbleSort(arr *[6]int) {
	// 外层循环控制排序的轮数，总共需要进行 n - 1 轮
	for i := 0; i < 6-1; i++ {
		// 标记是否发生交换，用于优化排序过程
		swapped := false
		// 内层循环用于比较相邻元素并交换位置
		for j := 0; j < 6-i-1; j++ {
			// 如果前一个元素大于后一个元素，则交换它们的位置
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				// 标记发生了交换
				swapped = true
			}
		}
		// 如果在一轮比较中没有发生交换，说明数组已经有序，可以提前结束排序
		if !swapped {
			break
		}
	}
}
