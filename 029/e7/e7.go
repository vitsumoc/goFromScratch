package main

import (
	"errors"
	"fmt"
)

func main() {
	var err error
	if true {
		i, err := errFun()
		fmt.Println(i)
		fmt.Println(err)
	}
	fmt.Println(err)
}

func errFun() (int, error) {
	return 1, errors.New("报错")
}
