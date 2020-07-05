package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	defer func() {
		fmt.Println("Foo inner deferred function running")
	}()
	fmt.Println("This is foo running")
}

func bar() {
	fmt.Println("This is bar running")
}
