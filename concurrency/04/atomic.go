package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	const times = 100

	var incrementer int64
	var wg sync.WaitGroup

	wg.Add(times)

	for x := 0; x < times; x++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			fmt.Println(atomic.LoadInt64(&incrementer))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final value:", incrementer)
}
