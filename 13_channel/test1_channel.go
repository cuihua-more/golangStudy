package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义一个channel
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutinue 结束")

		fmt.Println("goroutinue 正在运行...")

		// time.Sleep(1 * time.Second)
		c <- 66 // 把66值发给channel 写完以后，如果没有人读，则阻塞
	}()

	num := <-c // 从channle中读取内容，默认行为阻塞，等待channel中有内容
	fmt.Println("num = ", num)
	fmt.Println("main gorontine 结束")
}
