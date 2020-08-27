package main

import (
	"testing"
)

func TestCountEvenDigits(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{12, 345, 2, 6, 7896}, 2},
		{[]int{555, 901, 482, 1771}, 1},
		{[]int{1, 100, 10000}, 0},
		{[]int{10, 1000, 100000}, 3},
		{[]int{1, 10, 100, 1000, 10000, 100000}, 3},
	}

	for _, tt := range tests {
		got := countEvenDigits(tt.nums)
		if got != tt.want {
			t.Errorf("failed for %+v: got=%d\n", tt, got)
		}
	}
}
