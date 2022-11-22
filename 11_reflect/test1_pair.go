package main

import "fmt"

func main() {
	var a string

	// pair<static type: string, value: "abcd">
	a = "abcd"

	var allType interface{}

	// pair< type: string, value: "abcd">
	allType = a

	value, _ := allType.(string)
	fmt.Println(value)
}
