package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	n1    []int
	n1Len int
	n2    []int
	n2Len int
	want  []int
}

func TestMergeNaive(t *testing.T) {
	tests := []TestCase{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{}, 0, []int{1, 2, 3, 0, 0, 0}},
		{[]int{2, 0}, 1, []int{1}, 1, []int{1, 2}},
		{[]int{0, 0, 0}, 0, []int{3, 6, 9}, 3, []int{3, 6, 9}},
	}

	for _, tt := range tests {
		// Copy input as merge() merges in-place.
		nums := make([]int, len(tt.n1))
		copy(nums, tt.n1)

		mergeNaive(nums, tt.n1Len, tt.n2, tt.n2Len)
		if !reflect.DeepEqual(nums, tt.want) {
			t.Errorf("failed for %+v: got=%+v", tt, nums)
		}
	}
}

func TestMergeTwoPtrs(t *testing.T) {
	tests := []TestCase{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{}, 0, []int{1, 2, 3, 0, 0, 0}},
		{[]int{7, 8, 9, 0, 0, 0}, 3, []int{1, 2, 3}, 3, []int{1, 2, 3, 7, 8, 9}},
		{[]int{2, 0}, 1, []int{1}, 1, []int{1, 2}},
		{[]int{0, 0, 0}, 0, []int{3, 6, 9}, 3, []int{3, 6, 9}},
	}

	for _, tt := range tests {
		// Copy input as merge() merges in-place.
		nums := make([]int, len(tt.n1))
		copy(nums, tt.n1)

		mergeTwoPtrs(nums, tt.n1Len, tt.n2, tt.n2Len)
		if !reflect.DeepEqual(nums, tt.want) {
			t.Errorf("failed for %+v: got=%+v", tt, nums)
		}
	}
}
