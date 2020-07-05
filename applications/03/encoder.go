package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	First   string
	Last    string
	Age     int
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
	err := json.NewEncoder(os.Stdout).Encode(people)
	if err != nil {
		fmt.Println(err)
	}
}
