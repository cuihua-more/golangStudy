package main

import "fmt"

func main() {
	slice := []int{11, 22, 33, 44}

	fmt.Printf("len = %d, slice = %v\n", len(slice), slice) // %v 打印全部内容

	var slice1 []int        //声明一个slice1是一个切片，但是没有空间
	slice1 = make([]int, 3) // 分配空间

	slice1[1] = 100
	fmt.Printf("len = %d, slice1 = %v\n", len(slice1), slice1)

	var slice2 []int = make([]int, 5) // 声明切片，同时分配空间 var可省略
	for index, _ := range slice2 {
		slice2[index] = index
	}
	fmt.Printf("len = %d, slice2 = %v\n", len(slice2), slice2)

	//判断一个切片是否由空间
	var slice3 []int
	if slice3 == nil {
		fmt.Printf("slice3 is empty!\n")
	} else {
		fmt.Printf("slice3 is full!\n")
	}
}
