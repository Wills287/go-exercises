package main

import "fmt"

func main() {
	f := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		} else if len(xi) == 1 {
			return xi[0]
		}
		return xi[0] + xi[len(xi)-1]
	}
	x := doubler(f, []int{2, 3, 4, 5, 6})
	fmt.Println(x)
}

func doubler(f func(xi []int) int, ii []int) int {
	x := f(ii)
	return x * 2
}
