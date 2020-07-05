package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	for _, x := range []struct {
		input  []int
		output float64
		error  error
	}{
		{[]int{2, 3, 4, 5, 6, 7}, 4.5, nil},
		{[]int{2, 3, 4}, 3, nil},
		{[]int{1, 1, 2, 3, 5, 8, 13, 21}, 5.333333333333333, nil},
		{[]int{}, 0, insufficientSliceSize},
		{[]int{2, 3}, 0, insufficientSliceSize},
	} {
		n, err := CenteredAvg(x.input)
		if x.error != nil && x.error != err {
			t.Errorf("Expected %v, got %v", x.error, err)
		} else if n != x.output {
			t.Errorf("Expected %v, got %v", x.output, n)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{2, 3, 4, 5, 6, 7}))
	// Output:
	// 4.5 <nil>
}

func BenchmarkCenteredAvg(b *testing.B) {
	for x := 0; x < b.N; x++ {
		_, _ = CenteredAvg([]int{2, 3, 4, 5, 6, 7})
	}
}
