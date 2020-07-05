package main

import "fmt"

func main() {
	x := struct {
		firstName string
		lastName  string
		pets      map[string]string
		cars      []string
	}{
		firstName: "Test",
		lastName:  "Person",
		pets: map[string]string{
			"Doggie": "Dog",
			"Kitty":  "Cat",
		},
		cars: []string{"Reliant Robin", "Peel P50"},
	}
	fmt.Println(x.firstName)
	fmt.Println(x.lastName)
	for k, v := range x.pets {
		fmt.Println(k, ":", v)
	}
	for i, v := range x.cars {
		fmt.Println(i, ":", v)
	}
}
