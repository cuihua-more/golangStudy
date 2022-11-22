package main

import (
	"fmt"
	"reflect"
)

// reflect，反射包，就是定义一个元素，将元素的type和value得出
func reflectNum(arg interface{}) {
	fmt.Println("type:", reflect.TypeOf(arg))
	fmt.Println("value:", reflect.ValueOf(arg))
}

func main() {
	var num float64 = 1.23456

	reflectNum(num)
}
