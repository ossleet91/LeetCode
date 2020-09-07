package main

import (
	"reflect"
	"testing"
)

func TestMoveZeros(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 0, 0, 0, 4, 5, 0, 7}, []int{1, 2, 4, 5, 7, 0, 0, 0, 0}},
		{[]int{0, 0, 0, 0, 1, 4}, []int{1, 4, 0, 0, 0, 0}},
		{[]int{1, 2, 3, 6, 5, 4}, []int{1, 2, 3, 6, 5, 4}},
	}

	for _, tt := range tests {
		nums := make([]int, len(tt.nums))
		copy(nums, tt.nums)
		moveZeros(nums)
		if !reflect.DeepEqual(nums, tt.want) {
			t.Errorf("failed for %+v: got=%+v\n", tt, nums)
		}
	}
}

func TestMoveZerosLomuto(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 0, 0, 0, 4, 5, 0, 7}, []int{1, 2, 4, 5, 7, 0, 0, 0, 0}},
		{[]int{0, 0, 0, 0, 1, 4}, []int{1, 4, 0, 0, 0, 0}},
		{[]int{1, 2, 3, 6, 5, 4}, []int{1, 2, 3, 6, 5, 4}},
	}

	for _, tt := range tests {
		nums := make([]int, len(tt.nums))
		copy(nums, tt.nums)
		moveZerosLomuto(nums)
		if !reflect.DeepEqual(nums, tt.want) {
			t.Errorf("failed for %+v: got=%+v\n", tt, nums)
		}
	}
}
