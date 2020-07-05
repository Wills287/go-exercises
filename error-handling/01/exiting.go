package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p := person{
		First:   "Test",
		Last:    "Person",
		Sayings: []string{"Saying one", "Second saying"},
	}

	xb, err := json.Marshal(p)
	if err != nil {
		log.Fatalln("Unable to marshal json:", err)
	}
	fmt.Println(string(xb))
}
