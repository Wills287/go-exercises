package main

import "fmt"

func main() {
	x := wrapper()
	fmt.Println(x())
}

func wrapper() func() string {
	return func() string {
		return "Test"
	}
}
