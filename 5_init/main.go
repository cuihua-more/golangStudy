package main

import (
	_ "GolangStudy/5_init/lib1" // 匿名导入，只导入不使用
	// "GolangStudy/5_init/lib2"
	//. "GolangStudy/5_init/lib2" // 将包导入当前main包中
	mylib2 "GolangStudy/5_init/lib2" //别名导入
	"fmt"
)

func main() {
	fmt.Println("main.")

	// lib1.LibTest()
	mylib2.LibTest()
}
