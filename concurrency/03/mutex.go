package main

import (
	"fmt"
	"sync"
)

func main() {
	const times = 100

	var incrementer int
	var wg sync.WaitGroup
	var m sync.Mutex

	wg.Add(times)

	for x := 0; x < times; x++ {
		go func() {
			m.Lock()
			v := incrementer
			v++
			incrementer = v
			fmt.Println(incrementer)
			m.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final value:", incrementer)
}
