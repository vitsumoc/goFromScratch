package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(1 + 1)
	fmt.Println(reflect.TypeOf(1 + 1))
	fmt.Println(5 - 3)
	fmt.Println(reflect.TypeOf(5 - 3))
	fmt.Println(2 * 3)
	fmt.Println(reflect.TypeOf(2 * 3))
	fmt.Println(7 / 3)
	fmt.Println(reflect.TypeOf(7 / 3))
	fmt.Println(7 % 3)
	fmt.Println(reflect.TypeOf(7 % 3))
}
