package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(1.2 % 0.5)
	fmt.Println(reflect.TypeOf(1.2 % 0.5))
}
