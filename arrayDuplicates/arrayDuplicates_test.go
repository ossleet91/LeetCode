package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	nums []int
	want []int
}

func TestFindDuplicatesSpace(t *testing.T) {
	tests := []TestCase{
		{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{2, 3}},
	}

	for _, tt := range tests {
		nums := make([]int, len(tt.nums))
		copy(nums, tt.nums)

		got := findDuplicatesSpace(nums)

		if len(tt.want) == 0 {
			if len(got) != 0 {
				t.Errorf("failed for %+v due to bad gotLen: got=%v",
					tt, got)
			}
			continue
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("failed for %+v: got=%v\n", tt, got)
		}
	}
}

func TestFindDuplicates(t *testing.T) {
	tests := []TestCase{
		{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{2, 3}},
	}

	for _, tt := range tests {
		nums := make([]int, len(tt.nums))
		copy(nums, tt.nums)

		got := findDuplicates(nums)

		if len(tt.want) == 0 {
			if len(got) != 0 {
				t.Errorf("failed for %+v due to bad gotLen: got=%v",
					tt, got)
			}
			continue
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("failed for %+v: got=%v\n", tt, got)
		}
	}
}
