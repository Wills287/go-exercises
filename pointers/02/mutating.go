package main

import "fmt"

type person struct {
	first string
	last  string
}

func updatePerson(p *person) {
	p.first = "Changed"
	p.last = "Changed"
}

func main() {
	p := person{
		first: "First",
		last:  "Last",
	}
	fmt.Println(p)

	updatePerson(&p)
	fmt.Println(p)
}
