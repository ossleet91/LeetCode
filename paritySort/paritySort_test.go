package main

import (
	"testing"
)

type ParitySort struct {
	arr      []int // array to sort
	oddIndex int   // index at which odd begins; len(arr) represents no odd
}

func isAllEven(arr []int) bool {
	for _, n := range arr {
		if n%2 != 0 {
			return false
		}
	}

	return true
}

func isAllOdd(arr []int) bool {
	for _, n := range arr {
		if n%2 == 0 {
			return false
		}
	}

	return true
}

func TestParitySort(t *testing.T) {
	tests := []ParitySort{
		{[]int{3, 1, 2, 4}, 2},
		{[]int{55, 45, 36, 28, 21, 20, 15, 6, 3, 2}, 5},
		{[]int{0, 2, 4, 6, 8}, 5},
		{[]int{1, 3, 5}, 0},
	}

	for _, tt := range tests {
		got := paritySort(tt.arr)
		if !isAllEven(got[:tt.oddIndex]) {
			t.Errorf("odd mixed with even for %+v: got=%+v\n",
				tt, got)
		}
		if !isAllOdd(got[tt.oddIndex:]) {
			t.Errorf("even mixed with odd for %+v: got=%+v\n",
				tt, got)
		}
	}
}

func TestParitySortTwoPtrs(t *testing.T) {
	tests := []ParitySort{
		{[]int{3, 1, 2, 4}, 2},
		{[]int{55, 45, 36, 28, 21, 20, 15, 6, 3, 2}, 5},
		{[]int{0, 2, 4, 6, 8}, 5},
		{[]int{1, 3, 5}, 0},
	}

	for _, tt := range tests {
		got := paritySortTwoPtrs(tt.arr)
		if !isAllEven(got[:tt.oddIndex]) {
			t.Errorf("odd mixed with even for %+v: got=%+v\n",
				tt, got)
		}
		if !isAllOdd(got[tt.oddIndex:]) {
			t.Errorf("even mixed with odd for %+v: got=%+v\n",
				tt, got)
		}
	}
}
