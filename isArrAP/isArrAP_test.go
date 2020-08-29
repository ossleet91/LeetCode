package main

import (
	"testing"
)

type Test struct {
	arr  []int
	want bool
}

func TestCanMakeAPSort(t *testing.T) {
	tests := []Test{
		{[]int{3, 5, 1}, true},
		{[]int{1, 2, 4}, false},
	}

	for _, tt := range tests {
		got := canMakeAPSort(tt.arr)
		if got != tt.want {
			t.Errorf("failed for %+v: got=%t\n", tt, got)
		}
	}
}

func TestCanMakeAP(t *testing.T) {
	tests := []Test{
		{[]int{3, 5, 1}, true},
		{[]int{1, 2, 4}, false},
		{[]int{33, 21, 31, 19, 29, 23, 25, 27, 17}, true},
	}

	for _, tt := range tests {
		got := canMakeAP(tt.arr)
		if got != tt.want {
			t.Errorf("failed for %+v: got=%t\n", tt, got)
		}
	}
}
