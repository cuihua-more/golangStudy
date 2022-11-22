package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this *User) Call() {
	fmt.Println("user is called..")
	fmt.Printf("%v\n", this)
}

func main() {
	user := User{1, "Aced", 48}

	DoFiledAndMethod(user)
}

func DoFiledAndMethod(input interface{}) {
	// 获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is :", inputType.Name())

	// 获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue = ", inputValue)

	// 通过type获取里面的字段，也叫filed
	// 1. 通过interface获取reflect.type， 通过Type的NumFiled方法，获取定义数据的数量，进行遍历
	// 2. 通过Filed方法获取每一个数据类型
	// 3. 通过interface获取reflect.value， 通过Value的Filed(i).Interface()方法获取对应的值
	for i := 0; i < inputType.NumField(); i++ {
		filed := inputType.Field(i)              // 获取每一个定义
		value := inputValue.Field(i).Interface() // 获取对应的值

		fmt.Printf("%s: %v = %v\n", filed.Name, filed.Type, value)
	}

	//通过type获取里面的方法，并调用
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}
