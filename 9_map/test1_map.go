package main

import "fmt"

func main() {
	// 声明map类型，key是string类型， value是string，默认没有空间
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1 is empty!")
	}

	//分配10个键值对空间，当超过使用超过10个空间时，便会再分配10个，以此往复
	myMap1 = make(map[string]string, 10)

	myMap1["one"] = "C++"
	myMap1["two"] = "C"
	fmt.Println(myMap1)

	// 声明map第二种方式
	myMap2 := make(map[string]int) //默认用几个会分配一些空间，顺序不固定
	myMap2["one"] = 1
	myMap2["two"] = 2
	myMap2["three"] = 3
	fmt.Println(myMap2)
	fmt.Printf("len = %d, type is %T\n", len(myMap2), myMap2)

	//声明map的第三种方式
	myMap3 := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
	}
	fmt.Println(myMap3)
	fmt.Printf("len = %d, type is %T\n", len(myMap3), myMap3)

}
