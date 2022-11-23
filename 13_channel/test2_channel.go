package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3) //带缓冲的channel
	fmt.Println("len(c) = ", len(c), " cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("gotoutine 结束")

		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子goroutine is running, send value = ", i, "len(c)", len(c), " cap(c)", cap(c))
		}
	}()

	time.Sleep(1 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("num = ", num)
	}
	fmt.Println("main goroutine 结束")
}
