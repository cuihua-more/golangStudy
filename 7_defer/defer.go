package main

import "fmt"

/*
* defer，注意一： 多个defer，采用先进后出的顺序执行
 */

func main() {
	// defer关键字，表示这一行在函数最后执行，
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")

	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2")
}
