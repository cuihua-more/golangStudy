package main

import "fmt"

// const 用来定义枚举 iota在第一行为0，逐行+1，iota只能在const中定义
const (
	BEIJING  = iota //  0
	SHANGHAI        //  1
	SHENZHEN        //  2
)

const (
	aa = 10 * iota // 0
	bb             // 10
	cc             // 20
)

const (
	a, b = iota + 1, iota + 2 // a = 1 b = 2
	c, d                      // c = 2 d = 3
	e, f                      // e = 3 f = 4

	g, h = iota * 2, iota * 3 // g = 6, h = 9
	i, k                      // i = 8. k = 12
)

func main() {
	// 定义常量，将var 改为const
	const length int = 100

	fmt.Println("length = ", length)
}
