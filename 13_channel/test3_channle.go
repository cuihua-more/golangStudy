package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c) //关闭channel
	}()

	// for {
	// 	// ok表示channel有没有关闭，true表示开启；
	// 	// 下面if语句，先执行 data, ok := <-c; 再根据ok真假进行判断
	// 	if data, ok := <-c; ok {
	// 		fmt.Println(data)
	// 	} else {
	// 		break
	// 	}
	// }

	// range可以不断的操作channel
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("Main Finished...")
}
