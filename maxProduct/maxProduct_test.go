package main

import (
	"testing"
)

func TestMaxIndex(t *testing.T) {
	tests := []struct {
		nums      []int
		wantIndex int
		wantMax   int
	}{
		{[]int{3, 4, 5, 2}, 2, 5},
		{[]int{1, 5, 4, 5}, 1, 5},
		{[]int{3, 7}, 1, 7},
	}

	for _, tt := range tests {
		gotIndex, gotMax := maxIndex(tt.nums)
		if gotIndex != tt.wantIndex || gotMax != tt.wantMax {
			t.Errorf("failed for %+v: gotIndex=%d gotMax=%d\n",
				tt, gotIndex, gotMax)
		}
	}
}

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 4, 5, 2}, 12},
		{[]int{1, 5, 4, 5}, 16},
		{[]int{3, 7}, 12},
	}

	for _, tt := range tests {
		got := maxProduct(tt.nums)
		if got != tt.want {
			t.Errorf("failed for %+v: got=%d\n", tt, got)
		}
	}
}
