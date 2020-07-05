package main

import "fmt"

func main() {
	x, y := 10, 10
	c := generate(x, y)

	for z := 0; z < x*y; z++ {
		fmt.Println(<-c)
	}

	fmt.Println("Exiting")
}

func generate(x, y int) <-chan int {
	c := make(chan int)
	for a := 0; a < x; a++ {
		go func() {
			for b := 0; b < y; b++ {
				c <- b
			}
		}()
	}
	return c
}
