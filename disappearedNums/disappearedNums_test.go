package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	nums []int
	want []int
}

func TestFindDisappearedNumbersSpace(t *testing.T) {
	tests := []TestCase{
		{[]int{1, 2, 3, 4}, []int{}},
		{[]int{1, 1, 4, 4}, []int{2, 3}},
		{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{5, 6}},
	}

	for _, tt := range tests {
		nums := make([]int, len(tt.nums))
		copy(nums, tt.nums)

		got := findDisappearedNumbersSpace(nums)

		if len(tt.want) == 0 {
			if len(got) != 0 {
				t.Errorf("bad gotLen for %+v: got=%d", tt, got)
			}
			continue
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("failed for %+v: got=%+v", tt, got)
		}
	}
}

func TestFindDisappearedNumbers(t *testing.T) {
	tests := []TestCase{
		{[]int{1, 2, 3, 4}, []int{}},
		{[]int{1, 1, 4, 4}, []int{2, 3}},
		{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{5, 6}},
	}

	for _, tt := range tests {
		nums := make([]int, len(tt.nums))
		copy(nums, tt.nums)

		got := findDisappearedNumbers(nums)

		if len(tt.want) == 0 {
			if len(got) != 0 {
				t.Errorf("bad gotLen for %+v: got=%d", tt, got)
			}
			continue
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("failed for %+v: got=%+v", tt, got)
		}
	}
}
