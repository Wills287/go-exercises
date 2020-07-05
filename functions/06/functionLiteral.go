package main

import "fmt"

func main() {
	func() {
		for x := 1; x <= 3; x++ {
			fmt.Println(x)
		}
	}()

	fmt.Printf("Returned value from function: %v\n",
		func(s string) string {
			return fmt.Sprintf("---%v---", s)
		}("Testing"))
}
