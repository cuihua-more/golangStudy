package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	// 区间截取，s1和s共用一段内存，但是s1的长度只有0~2这么长，不算2
	s1 := s[0:2] // 相当于s[0], s[1]

	fmt.Println(s)
	fmt.Println(s1)

	s[0] = 200

	fmt.Println(s)
	fmt.Println(s1)

	fmt.Printf("--------------------\n")
	// copy函数拷贝内容
	s2 := make([]int, 4)
	copy(s2, s)
	fmt.Println(s2)
}
