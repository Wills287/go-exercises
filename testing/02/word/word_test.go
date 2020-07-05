package word

import (
	"../quote"
	"fmt"
	"testing"
)

func TestUsageCount(t *testing.T) {
	for _, x := range []struct {
		input  string
		output map[string]int
	}{
		{
			"One two three two three three",
			map[string]int{
				"one":   1,
				"two":   2,
				"three": 3,
			},
		},
		{
			"One t t t t t",
			map[string]int{
				"one": 1,
				"t":   5,
			},
		},
		{
			"One",
			map[string]int{
				"one": 1,
			},
		},
		{
			"Onetwothree",
			map[string]int{
				"onetwothree": 1,
			},
		},
		{
			"One      two three",
			map[string]int{
				"one":   1,
				"two":   1,
				"three": 1,
			},
		},
		{
			"One? one. one!",
			map[string]int{
				"one": 3,
			},
		},
		{
			"One...one...one!",
			map[string]int{
				"one": 3,
			},
		},
	} {
		m := UsageCount(x.input)
		for k, v := range m {
			if v != x.output[k] {
				t.Errorf("Expected {%v : %v}, got {%v : %v}", k, x.output[k], k, v)
			}
		}
	}
}

func TestTotalCount(t *testing.T) {
	for _, x := range []struct {
		input  string
		output int
	}{
		{"One two three two three three", 6},
		{"One t t t t t", 6},
		{"One", 1},
		{"Onetwothree", 1},
		{"One      two three", 3},
		{"One? one. one!", 3},
		{"One...one...one!", 3},
	} {
		n := TotalCount(x.input)
		if n != x.output {
			t.Errorf("Expected %v, got %v", x.output, n)
		}
	}
}

func ExampleUsageCount() {
	fmt.Println(UsageCount("One two three two three three"))
	// Output:
	// map[one:1 three:3 two:2]
}

func ExampleTotalCount() {
	fmt.Println(TotalCount("One two three two three three"))
	// Output:
	// 6
}

func BenchmarkTotalCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TotalCount(quote.SunAlso)
	}
}

func BenchmarkUsageCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UsageCount(quote.SunAlso)
	}
}
