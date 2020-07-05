package main

import "fmt"

type customError struct {
	info string
}

func (ce customError) Error() string {
	return fmt.Sprintf("an error occurred: %v", ce.info)
}

func main() {
	ce := customError{
		info: "need more coffee",
	}
	printError(ce)
}

func printError(e error) {
	fmt.Printf("Printing error - %v\n", e)
	fmt.Printf("Printing info - %v\n", e.(customError).info)
}
