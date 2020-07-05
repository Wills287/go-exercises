package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	go func() {
		q <- 1
	}()
	c := generate(q)

	receive(c, q)

	fmt.Println("Exiting")
}

func generate(q chan int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		//q <- 1
		close(c)
	}()

	return c
}

func receive(c <-chan int, q chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			fmt.Println("Quit select")
			return
		}
	}
}
