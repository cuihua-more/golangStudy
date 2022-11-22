package main

import "fmt"

// interface{} 这种类型，是一个指针，类似与C语言的void *指针
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	// Golang给interface{}提供了一种变量类型断言机制
	value, ok := arg.(string) //这个方式断言arg是否是string类型
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type, value = ", value)

		fmt.Printf("value type is %T\n", value)
	}
}

type Book struct {
	auth string
}

func main() {
	book := Book{"Golang"}
	myFunc(book)
	myFunc(100)
	myFunc("123")
	myFunc(3.14)
}
