package main

import (
	"reflect"
	"testing"
)

func TestShift(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{1, 0, 2, 3, 0, 4, 5, 0}, []int{1, 1, 0, 2, 3, 0, 4, 5}},
		{[]int{1, 2, 3}, []int{1, 1, 2}},
		{[]int{1}, []int{1}},
		{[]int{-50, -25}, []int{-50, -50}},
	}

	for _, tt := range tests {
		arr := make([]int, len(tt.arr))
		copy(arr, tt.arr)

		shift(arr)
		if !reflect.DeepEqual(arr, tt.want) {
			t.Errorf("failed for %+v: got=%+v\n", tt, arr)
		}
	}
}

type TestCase struct {
	arr  []int
	want []int
}

func TestDuplicateZerosNaive(t *testing.T) {
	tests := []TestCase{
		{[]int{1, 0, 2, 3, 0, 4, 5, 0}, []int{1, 0, 0, 2, 3, 0, 0, 4}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2, 0, 3, 0, 4}, []int{1, 2, 0, 0, 3, 0}},
	}

	for _, tt := range tests {
		arr := make([]int, len(tt.arr))
		copy(arr, tt.arr)

		duplicateZerosNaive(arr)
		if !reflect.DeepEqual(arr, tt.want) {
			t.Errorf("failed for %+v: got=%+v\n", tt, arr)
		}
	}
}

func TestDuplicateZerosTwoPtrs(t *testing.T) {
	tests := []TestCase{
		{[]int{1, 0, 2, 3, 0, 4, 5, 0}, []int{1, 0, 0, 2, 3, 0, 0, 4}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2, 0, 3, 0, 4}, []int{1, 2, 0, 0, 3, 0}},
	}

	for _, tt := range tests {
		arr := make([]int, len(tt.arr))
		copy(arr, tt.arr)

		duplicateZerosTwoPtrs(arr)
		if !reflect.DeepEqual(arr, tt.want) {
			t.Errorf("failed for %+v: got=%+v\n", tt, arr)
		}
	}
}
