package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"Name" doc:"Doc"` // 这个是结构体的标签，键值对的方式，可以写多个
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem() // 获取结构体类型

	for i := 0; i < t.NumField(); i++ {
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info :", taginfo, " doc:", tagdoc)
	}

}

func main() {
	var re resume
	findTag(&re)
}
