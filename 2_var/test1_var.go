package main

import "fmt"

// 生命全局变量，方法一、二、三是可以的，方法四不可以
var g_a = "global var"

func main() {
	// 定义变量方式一，手动定义类型，默认值为0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a) //%T ,表示变量类型，int, string......

	// 定义变量方式二，手动定义类型，手动幅初值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	// 定义变量方式三，根据赋初值自适应变量类型
	var c = 1000
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	var cc = "aabbcc"
	fmt.Println("cc = ", cc)
	fmt.Printf("type of cc = %T\n", cc)

	// 定义变量方法四，省略var，根据初值自动适应变量类型，通过 :
	// 只能用于函数内
	e := 1000
	ee := "aabbccdd"
	eee := 3.14
	fmt.Printf("e = %d, type of e = %T\n", e, e)
	fmt.Printf("ee = %s, type of ee = %T\n", ee, ee)
	fmt.Printf("eee = %.2f, type of eee = %T\n", eee, eee)

	// 全局变量
	fmt.Printf("g_a = %s, type of g_a = %T\n", g_a, g_a)

	/*
	 * 单行声明多个变量
	 */
	var x, y int = 100, 200
	fmt.Printf("x = %d, y = %d\n", x, y)
	var xx, yy = 300, "string_var"
	fmt.Println("xx = ", xx, " yy = ", yy)

	/*
	* 多行声明多个变量
	 */
	var (
		vv int    = 100
		nn string = "string_vat"
		bb bool   = true
	)
	fmt.Println("vv = ", vv, " nn = ", nn, "bb = ", bb)
}
