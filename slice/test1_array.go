package main

import "fmt"

// 固定长度数组传参，传递方式为值拷贝
func printArrag(myArray [4]int) {
	fmt.Printf("run printArrag...\n")
	for index, value := range myArray {
		fmt.Println("index = ", index, " value = ", value)
	}
}

func main() {
	// 固定数组
	var myArrag [10]int

	myArray2 := [10]int{1, 2, 3, 4}

	for i := 0; i < len(myArrag); i++ {
		fmt.Println(myArrag[i])
	}

	// for循环另一种写法
	for index, value := range myArray2 {
		fmt.Println("index = ", index, " value = ", value)
	}

	// 查看数组的数据类型
	fmt.Printf("myArray1 types = %T\n", myArrag)
	fmt.Printf("myArray2 types = %T\n", myArray2)

	myArray3 := [4]int{11, 22, 33, 44}
	printArrag(myArray3)
}
