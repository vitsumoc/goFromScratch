package main

import "fmt"

func main() {
	var mac int = 0x68DA73A2B954

	fmt.Printf("MAC地址的十六进制表示是：0x%x\n", mac)
	fmt.Printf("MAC地址的十进制表示是：%d\n", mac)
	fmt.Printf("MAC地址的二进制表示是：0b%b\n", mac)
}
