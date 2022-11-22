package main

import "fmt"

// 声明一种数据类型，和C语言，typedef unsigned int uint32;一样
type myint int

// 定义一个结构体
type Book struct {
	title string
	auth  string
}

// 结构体传参是内存拷贝，不是传递地址
func changBook(book Book) {
	fmt.Printf("run changeBook...")
	fmt.Println("title = ", book.title, " auth = ", book.auth)

	book.auth = "li4"
}

// 可以传指针
func changBook2(book *Book) {
	fmt.Printf("run changeBook2...")
	fmt.Println("title = ", book.title, " auth = ", book.auth)

	book.auth = "111"
}

func main() {
	var a myint = 10
	fmt.Printf("a = %d, type is %T\n", a, a)

	var book1 Book
	book1.title = "Golang"
	book1.auth = "zhang3"

	fmt.Printf("%v\n", book1)

	fmt.Println("--------------------")
	changBook(book1)
	fmt.Printf("%v\n", book1)

	fmt.Println("--------------------")
	changBook2(&book1)
	fmt.Printf("%v\n", book1)

}
