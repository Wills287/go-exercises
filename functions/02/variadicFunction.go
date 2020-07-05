package main

import "fmt"

func main() {
	xi := []int{3, 4, 5, 6, 7}

	fmt.Println(foo(xi...))
	fmt.Println(bar(xi))
}

func foo(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func bar(xi []int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
