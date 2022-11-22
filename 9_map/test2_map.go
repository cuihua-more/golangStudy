package main

import (
	"fmt"
)

func printMap(myMap map[string]string) {
	fmt.Println("run printMap...")

	for key, value := range myMap {
		fmt.Println("key = ", key, " value = ", value)
		if key == "Chinese" {
			myMap[key] = "Henan"
		}
	}

}

func main() {
	cityMap := make(map[string]string)

	// 添加key value
	cityMap["Chinese"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	// 遍历
	for key, value := range cityMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}

	//删除
	delete(cityMap, "Japan")

	fmt.Println("----------------------------")
	for key, value := range cityMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}

	// 修改
	cityMap["USA"] = "DC"
	fmt.Println("----------------------------")
	for key, value := range cityMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}

	// 函数参数传递map,指针传递
	fmt.Println("----------------------------")
	printMap(cityMap)

	fmt.Println("----------------------------")
	for key, value := range cityMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}
}
