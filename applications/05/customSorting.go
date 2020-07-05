package main

import (
	"fmt"
	"sort"
)

type person struct {
	First string
	Last  string
	Age   int
}

type byLast []person
type byAge []person

func (x byLast) Len() int           { return len(x) }
func (x byLast) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x byLast) Less(i, j int) bool { return x[i].Last < x[j].Last }

func (x byAge) Len() int           { return len(x) }
func (x byAge) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x byAge) Less(i, j int) bool { return x[i].Age < x[j].Age }

func main() {
	people := []person{
		{
			First: "First",
			Last:  "Name",
			Age:   33,
		}, {
			First: "Second",
			Last:  "Person",
			Age:   28,
		}, {
			First: "Third",
			Last:  "Hello",
			Age:   25,
		},
	}
	fmt.Println("Unsorted", people)
	sort.Sort(byLast(people))
	fmt.Println("ByLast", people)

	sort.Sort(byAge(people))
	fmt.Println("ByAge", people)
}
