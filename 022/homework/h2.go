package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) Greet() string {
	return fmt.Sprintf("你好，我是%s，今年%d岁\n", p.Name, p.Age)
}

func main() {
	yiXing := Person{
		Name: "张一猩",
		Age:  21,
	}
	fmt.Printf(yiXing.Greet())
}
