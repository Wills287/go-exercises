package main

import (
	"./quote"
	"./word"
	"fmt"
)

func main() {
	fmt.Println(word.TotalCount(quote.SunAlso))

	for k, v := range word.UsageCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
