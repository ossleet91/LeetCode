package main

import "testing"

type TestCase struct {
	heights []int
	want    int
}

func TestHeightCheckerSort(t *testing.T) {
	tests := []TestCase{
		{[]int{1, 1, 4, 2, 1, 3}, 3},
		{[]int{5, 1, 2, 3, 4}, 5},
		{[]int{1, 2, 3, 4, 5}, 0},
	}

	for _, tt := range tests {
		got := heightCheckerSort(tt.heights)
		if got != tt.want {
			t.Errorf("failed for %+v: got=%d", tt, got)
		}
	}
}
