package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "汪汪汪"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "喵喵喵"
}

func main() {
	fmt.Printf("狗狗的叫声是：%s\n", Dog{}.Speak())
	fmt.Printf("猫猫的叫声是：%s\n", Cat{}.Speak())
}
