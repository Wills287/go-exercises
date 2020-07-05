package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	for _, x := range []struct {
		input  int
		output int
	}{
		{1, 7},
		{2, 14},
		{5, 35},
		{10, 70},
		{0, 0},
		{-1, -7},
	} {
		v := Years(x.input)
		if v != x.output {
			t.Errorf("Expected %v, got %v", x.output, v)
		}
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	// Output:
	// 70
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	// Output:
	// 70
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
