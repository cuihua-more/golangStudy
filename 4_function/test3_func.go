package main

import "fmt"

/*
* int fool(string a, int b)
 */
func foo1(a string, b int) int {
	fmt.Printf("------%s------\n", "foo1")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100

	return c
}

/*
* 返回两个值， 匿名
 */
func foo2(a string, b int) (int, bool) {
	fmt.Printf("------%s------\n", "foo2")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return 200, false
}

/*
* 返回两个值， 有形参
 */
func foo3(a string, b int) (r1 int, r2 bool) {
	fmt.Printf("------%s------\n", "foo3")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	r1 = 1000
	r2 = true
	return
}

/*
* 返回两个值， 有形参, 类型一样，可以只写一个变量
 */
func foo4(a string, b int) (r1, r2 int) {
	fmt.Printf("------%s------\n", "foo4")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	r1 = 50
	r2 = 100
	return
}

func main() {
	a := foo1("abc", 555)

	fmt.Println("a = ", a)

	b, c := foo2("def", 666)
	fmt.Println("b = ", b, "c = ", c)

	b, c = foo3("def", 666)
	fmt.Println("b = ", b, "c = ", c)

	d, e := foo4("def", 666)
	fmt.Println("b = ", d, "c = ", e)
}
