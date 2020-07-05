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

	xb, err := toJson(p)
	if err != nil {
		log.Fatalln("Unable to retrieve json:", err)
	}
	fmt.Println(string(xb))
}

func toJson(v interface{}) ([]byte, error) {
	xb, err := json.Marshal(v)
	if err == nil {
		return []byte{}, fmt.Errorf("error occurred marshalling json for %v: %w", v, err)
	}
	return xb, nil
}
