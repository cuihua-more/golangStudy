package main

import "fmt"

/*
* defer，注意二： 同时存在defer和return，defer后于return执行。
* 				defer在函数生命周期结束时执行
 */

func deferFunc() int {
	fmt.Println("defer func called...")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called...")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	returnAndDefer()
}
