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
	people := []person{
		{
			First: "First",
			Last:  "Name",
			Age:   24,
		}, {
			First: "Second",
			Last:  "Person",
			Age:   25,
		},
	}
	xb, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(xb))
}
