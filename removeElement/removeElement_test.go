package main

import (
	"testing"
)

type TestCase struct {
	nums    []int
	val     int
	wantLen int
}

// Returns true if arr contains val.
func contains(arr []int, val int) bool {
	for _, n := range arr {
		if n == val {
			return true
		}
	}
	return false
}

func TestRemoveElement(t *testing.T) {
	tests := []TestCase{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 2, 3, 4, 5, 3, 1, 0}, 3, 6},
		{[]int{1, 1, 1}, 1, 0},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
		{[]int{2}, 3, 1},
		{[]int{3, 3}, 5, 2},
	}

	for _, tt := range tests {
		nums := make([]int, len(tt.nums))
		copy(nums, tt.nums)

		gotLen := removeElement(nums, tt.val)
		if tt.wantLen != gotLen {
			t.Errorf("failed for %+v: bad length: gotLen=%d got=%+v\n", tt, gotLen, nums)
		}

		if contains(nums[:gotLen], tt.val) {
			t.Errorf("failed for %+v: bad array: gotLen=%d got=%+v\n", tt, gotLen, nums)
		}
	}
}

func TestRemoveElement1(t *testing.T) {
	tests := []TestCase{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 2, 3, 4, 5, 3, 1, 0}, 3, 6},
		{[]int{1, 1, 1}, 1, 0},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
		{[]int{2}, 3, 1},
		{[]int{3, 3}, 5, 2},
	}

	for _, tt := range tests {
		nums := make([]int, len(tt.nums))
		copy(nums, tt.nums)

		gotLen := removeElement1(nums, tt.val)
		if tt.wantLen != gotLen {
			t.Errorf("failed for %+v: bad length: gotLen=%d got=%+v\n", tt, gotLen, nums)
		}

		if contains(nums[:gotLen], tt.val) {
			t.Errorf("failed for %+v: bad array: gotLen=%d got=%+v\n", tt, gotLen, nums)
		}
	}
}

func TestRemoveElement2(t *testing.T) {
	tests := []TestCase{
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 2, 3, 4, 5, 3, 1, 0}, 3, 6},
		{[]int{1, 1, 1}, 1, 0},
		{[]int{2}, 3, 1},
		{[]int{3, 3}, 5, 2},
	}

	for _, tt := range tests {
		nums := make([]int, len(tt.nums))
		copy(nums, tt.nums)

		gotLen := removeElement2(nums, tt.val)
		if tt.wantLen != gotLen {
			t.Errorf("failed for %+v: bad length: gotLen=%d got=%+v\n", tt, gotLen, nums)
		}

		if contains(nums[:gotLen], tt.val) {
			t.Errorf("failed for %+v: bad array: gotLen=%d got=%+v\n", tt, gotLen, nums)
		}
	}
}
