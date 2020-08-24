package main

import (
	"testing"
)

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 2, 1}, 1},
	}

	for _, tt := range tests {
		got := singleNumber(tt.nums)
		if tt.want != got {
			t.Errorf("failed for %+v: got=%d\n", tt, got)
		}
	}
}
