package main

import (
	"testing"
)

func sum(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	return sum
}

func isUnique(nums []int) bool {
	unique := make(map[int]bool)
	for _, n := range nums {
		seen, ok := unique[n]
		if ok && seen {
			return false
		}
		unique[n] = true
	}

	return true
}

func TestZeroSum(t *testing.T) {
	tests := []int{5, 3, 1}

	for _, tt := range tests {
		got := zeroSum(tt)

		if len(got) != tt {
			t.Errorf("incorrect length: want=%d got=%d gotArray=%+v\n",
				tt, len(got), got)
		}

		if !isUnique(got) {
			t.Errorf("returned array has non-unique elements: gotArray=%+v\n",
				got)
		}

		gotSum := sum(got)
		if gotSum != 0 {
			t.Errorf("sum is not zero: want=0 got=%d gotArray=%+v\n",
				gotSum, got)
		}
	}
}
