package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	const times = 100

	var incrementer int
	var wg sync.WaitGroup

	wg.Add(times)

	for x := 0; x < times; x++ {
		go func() {
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final value:", incrementer)
}
