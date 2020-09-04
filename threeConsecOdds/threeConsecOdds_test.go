package main

import (
	"testing"
)

type TestCase struct {
	arr  []int
	want bool
}

func TestThreeConsecutiveOdds(t *testing.T) {
	tests := []TestCase{
		{[]int{2, 6, 4, 1}, false},
		{[]int{1, 2, 34, 3, 4, 5, 7, 23, 12}, true},
		{[]int{1, 1, 1}, true},
	}

	for _, tt := range tests {
		got := threeConsecutiveOdds(tt.arr)
		if got != tt.want {
			t.Errorf("failed for %#v: got=%t\n", tt, got)
		}
	}
}
