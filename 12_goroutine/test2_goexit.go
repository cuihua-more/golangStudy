package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	// 用go执行一个形参为空返回值为空的携程，
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			// 退出当前goroutinue
			runtime.Goexit()
			fmt.Println("B")
		}() // ()表示定义完匿名函数后，去执行
		fmt.Println("A")
	}()

	// go创建一个带参数的匿名携程，但是目前返回值得不到，需要用channel
	go func(a int, b int) bool {
		fmt.Println("a = ", a, " b = ", b)
		return true
	}(10, 20)

	for {
		time.Sleep(1 * time.Second)
	}
}
