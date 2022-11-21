package lib1

import "fmt"

// 函数名首字母大写的话，说明这个函数是这个包提供的API，可以由外部调用，小写不行
func LibTest() {
	fmt.Println("lib1 test.")
}

func init() {
	fmt.Println("lib1 init.")
}
