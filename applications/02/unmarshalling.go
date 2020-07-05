package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Age   int    `json:"age"`
}

func main() {
	s := "[{\"first\":\"First\",\"last\":\"Name\",\"age\":24},{\"first\":\"Second\",\"last\":\"Person\",\"age\":25}]"
	var people []person
	err := json.Unmarshal([]byte(s), &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)
}
