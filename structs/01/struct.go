package main

import "fmt"

type person struct {
	firstName         string
	lastName          string
	favouriteFlavours []string
}

func main() {
	one := person{
		firstName:         "Test",
		lastName:          "Person",
		favouriteFlavours: []string{"Vanilla", "Chocolate"},
	}
	two := person{
		firstName:         "Another",
		lastName:          "Somebody",
		favouriteFlavours: []string{"Strawberry"},
	}

	printPerson(one)
	printPerson(two)
}

func printPerson(p person) {
	fmt.Println(p.firstName)
	fmt.Println(p.lastName)
	for _, v := range p.favouriteFlavours {
		fmt.Println(v)
	}
}
