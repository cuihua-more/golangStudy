package main

import "fmt"

func main() {
	number := make([]int, 3, 5)                                                      // 定义一个切片，长度为3，容量为5,相当于申请了5个int空间, cap说明是每一次申请的空间
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(number), cap(number), number) // cap函数是容量

	number = append(number, 1)                                                       // 追加一个元素1 [0 0 0 1]
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(number), cap(number), number) // 4 5 [0 0 0 1]

	number = append(number, 4) // 追加一个元素5 [0 0 0 1 5]
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(number), cap(number), number)

	number = append(number, 33) //当前容量已满，所以会再申请cap个int空间
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(number), cap(number), number)

	number2 := make([]int, 4) //初始切片不写cap，默认cap = len
	fmt.Printf("len = %d, cap = %d, slice2 = %v\n", len(number2), cap(number2), number2)

	number2 = append(number2, 44) // append时，会默认已满，再申请cap个空间，
	fmt.Printf("len = %d, cap = %d, slice2 = %v\n", len(number2), cap(number2), number2)

}
