package main

import (
	"fmt"
	"time"
)

// 从 goroutinue
func newTask() {
	i := 0

	for {
		i++
		fmt.Printf("new Goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

// 主goroutinue
func main() {
	// 创建一个go协程，执行newTask
	go newTask()

	i := 0
	for {
		i++
		fmt.Println("main Goroutine: i = ", i)
		time.Sleep(1 * time.Second)
	}
}
