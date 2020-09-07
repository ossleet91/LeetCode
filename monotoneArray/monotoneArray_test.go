package main

import "testing"

type TestCase struct {
	arr  []int
	want bool
}

func TestIsMonotonic(t *testing.T) {
	tests := []TestCase{
		{[]int{1, 2, 2, 3}, true},
		{[]int{6, 5, 4, 4}, true},
		{[]int{6, 5, 8, 8}, false},
		{[]int{1, 3, 2}, false},
		{[]int{1, 2, 4, 5}, true},
		{[]int{1, 1, 1}, true},
		{[]int{5, 5, 5, 5, 5, 4, 3}, true},
		{[]int{5, 5, 5, 5, 5, 4, 5}, false},
	}

	for _, tt := range tests {
		got := isMonotonic(tt.arr)
		if got != tt.want {
			t.Errorf("failed for %#v: got=%t\n", tt, got)
		}
	}
}
