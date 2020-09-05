package main

import (
	"testing"
)

type TestCase struct {
	nums []int
	want int
}

func TestMajorityElement(t *testing.T) {
	tests := []TestCase{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 2, 2}, 2},
	}

	for _, tt := range tests {
		got := majorityElement(tt.nums)
		if got != tt.want {
			t.Errorf("failed for %#v: got=%d\n", tt, got)
		}
	}
}

func TestMajorityElementBoyerMoore(t *testing.T) {
	tests := []TestCase{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 2, 2}, 2},
	}

	for _, tt := range tests {
		got := majorityElementBoyerMoore(tt.nums)
		if got != tt.want {
			t.Errorf("failed for %#v: got=%d\n", tt, got)
		}
	}
}
