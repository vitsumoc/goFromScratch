package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf("HelloWorld"))
	fmt.Println(reflect.TypeOf("1"))
	fmt.Println(reflect.TypeOf(10))
	fmt.Println(reflect.TypeOf(-1))
	fmt.Println(reflect.TypeOf(3.14))
	fmt.Println(reflect.TypeOf(1.0))
}
