package main

import "fmt"

func main() {
	bools := [8]bool{true, true, true, false, false, false, true, false}
	fmt.Printf("[8]bool 的值是：%v\n", bools)
	var u uint8 = bools2uint8(bools)
	fmt.Printf("压缩后的值是 %v\n", u)
	fmt.Printf("压缩后的值的二进制表示是 %08b\n", u)
	bs := uint82bools(u)
	fmt.Printf("解压缩后的的值是：%v\n", bs)
}

func bools2uint8(bools [8]bool) uint8 {
	var i uint8 = 0
	for x := 0; x < 8; x++ {
		if bools[x] {
			i += 1
		}
		if x != 7 { // 最后一次不左移
			i <<= 1
		}
	}
	return i
}

func uint82bools(u uint8) [8]bool {
	var bools [8]bool
	for x := 0; x < 8; x++ {
		if u&1 == 1 {
			bools[8-x-1] = true
		} else {
			bools[8-x-1] = false
		}
		u >>= 1
	}
	return bools
}
