package mymath

import (
	"errors"
	"sort"
)

var insufficientSliceSize = errors.New("provided slice too small to find centered average")

func CenteredAvg(xi []int) (float64, error) {
	if len(xi) < 3 {
		return 0, insufficientSliceSize
	}

	sort.Ints(xi)
	xi = xi[1:(len(xi) - 1)]

	n := 0
	for _, v := range xi {
		n += v
	}

	f := float64(n) / float64(len(xi))
	return f, nil
}
