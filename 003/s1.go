package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello" + "World!")
	fmt.Println(reflect.TypeOf("Hello" + "World!"))
}
