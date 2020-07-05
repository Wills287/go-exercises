package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Function one!")
		wg.Done()
	}()
	go func() {
		fmt.Println("Function two!")
		wg.Done()
	}()

	fmt.Println("Before wait")
	wg.Wait()
	fmt.Println("After wait")
}
