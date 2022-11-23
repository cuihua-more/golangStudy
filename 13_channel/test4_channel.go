package main

import "fmt"

func fibonacii(c, quit chan int) {
	x, y := 1, 1

	for {
		// select监控case中的channel情况，当条件满足时，就开始执行
		select {
		case c <- x:
			// 如果c可写，就执行以下
			y = x + y
			x = y
		case <-quit:
			// 如果quit可读，就执行以下
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)

	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}

		quit <- 0
	}()

	fibonacii(c, quit)
}
