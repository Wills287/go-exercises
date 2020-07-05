package main

import "fmt"

func main() {
	looper := func() {
		for x := 1; x <= 3; x++ {
			fmt.Println(x)
		}
	}
	looper()
	looper()
}
