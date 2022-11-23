package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 20, []string{"zhouoxingchi", "zhangbozhi"}}

	// 编码 结构体 -> json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error: ", err)
		return
	}

	fmt.Printf("jsonStr = %s\n", jsonStr)

	// 解码 json -> 结构体
	// jsonStr = {"title":"喜剧之王","year":2000,"rmb":20,"actors":["zhouoxingchi","zhangbozhi"]}
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal error: ", err)
		return
	}

	fmt.Println(myMovie)
}
