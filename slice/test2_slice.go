package main

import "fmt"

// 动态数组传递方式，类似指针传地址
func printArray(myArray []int) {
	// _表示匿名变量

	for _, value := range myArray {
		fmt.Println("value = ", value)
	}

	myArray[2] = 1000
}

func main() {
	myArray := []int{1, 2, 3, 4} // 动态数组，切片 slice

	fmt.Printf("myArray types is %T \n", myArray)

	printArray(myArray)

	fmt.Println("------------------")

	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
}
