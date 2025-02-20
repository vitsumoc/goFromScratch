package main

import "fmt"

func main() {
	const human = 7
	const cat = human * 7
	const mouse = cat * 7
	const ear = mouse * 7
	const wheat = ear * 7
	const sum = human + cat + mouse + ear + wheat
	fmt.Println(sum)
}
