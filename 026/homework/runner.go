package main

import (
	"fmt"
	"time"
)

type Runner struct {
	name     string
	distance int
}

func main() {
	runner1 := Runner{name: "张三", distance: 0}
	runner2 := Runner{name: "李四", distance: 0}
	runner3 := Runner{name: "王五", distance: 0}

	go func() {
		for x := 0; x < 10; x++ {
			run(&runner1)
			time.Sleep(300 * time.Millisecond)
		}
	}()

	go func() {
		for x := 0; x < 10; x++ {
			run(&runner2)
			time.Sleep(300 * time.Millisecond)
		}
	}()

	go func() {
		for x := 0; x < 10; x++ {
			run(&runner3)
			time.Sleep(300 * time.Millisecond)
		}
	}()

	time.Sleep(5 * time.Second)
}

func run(runner *Runner) {
	runner.distance += 10
	fmt.Printf("%s 跑了 %d 米\n", runner.name, runner.distance)
}
