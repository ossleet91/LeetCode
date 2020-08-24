package main

import (
	"testing"
)

func TestNumIdenticalPairs(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 1, 1, 3}, 4},
		{[]int{1, 1, 1, 1}, 6},
		{[]int{1, 2, 3}, 0},
	}

	for _, tt := range tests {
		got := numIdenticalPairs(tt.nums)
		if tt.want != got {
			t.Errorf("failed for %+v: got=%d\n", tt, got)
		}
	}
}
