package main

import "fmt"

type person struct {
	firstName         string
	lastName          string
	favouriteFlavours []string
}

func main() {
	m := map[string]person{
		"Person": {
			firstName:         "Test",
			lastName:          "Person",
			favouriteFlavours: []string{"Vanilla", "Chocolate"},
		},
		"Somebody": {
			firstName:         "Another",
			lastName:          "Somebody",
			favouriteFlavours: []string{"Strawberry"},
		},
	}

	for _, v := range m {
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for _, f := range v.favouriteFlavours {
			fmt.Println(f)
		}
		fmt.Println("----------")
	}
}
