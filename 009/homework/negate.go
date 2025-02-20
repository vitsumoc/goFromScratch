package main

import "fmt"

func main() {
	var i1 int = 1
	var i2 int = 35
	var i3 int = -9998877
	negate(&i1)
	negate(&i2)
	negate(&i3)
	fmt.Printf("1 取反后的值为 %d\n", i1)
	fmt.Printf("35 取反后的值为 %d\n", i2)
	fmt.Printf("-9998877 取反后的值为 %d", i3)
}

func negate(i *int) {
	*i = 0 - *i
}
