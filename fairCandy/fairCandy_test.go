package main

import (
	"testing"
)

type TestCase struct {
	alice []int
	bob   []int
}

func TestFairCandySwap(t *testing.T) {
	tests := []TestCase{
		{[]int{1, 1}, []int{2, 2}},
		{[]int{1, 2}, []int{2, 3}},
		{[]int{2}, []int{1, 3}},
		{[]int{1, 2, 5}, []int{2, 4}},
	}

	for _, tt := range tests {
		got := fairCandySwap(tt.alice, tt.bob)
		ashare, bshare := got[0], got[1]
		asum, bsum := sum(tt.alice), sum(tt.bob)
		if asum-ashare+bshare != bsum-bshare+ashare {
			t.Errorf("failed for %#v: got=%#v\n", tt, got)
		}
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 1}, 2},
		{[]int{1, 2, 5}, 8},
		{[]int{0, 0, 0}, 0},
	}

	for _, tt := range tests {
		got := sum(tt.nums)
		if got != tt.want {
			t.Errorf("failed for %#v: got=%d\n", tt, got)
		}
	}
}
