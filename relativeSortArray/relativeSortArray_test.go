package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	arr1 []int
	arr2 []int
	want []int
}

func TestRelativeSortArray(t *testing.T) {
	tests := []TestCase{
		{
			[]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
			[]int{2, 1, 4, 3, 9, 6},
			[]int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
		},
	}

	for _, tt := range tests {
		got := relativeSortArray(tt.arr1, tt.arr2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("failed for %#v: got=%#v\n", tt, got)
		}
	}
}

func TestRelativeSortArrayLessSpace(t *testing.T) {
	tests := []TestCase{
		{
			[]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
			[]int{2, 1, 4, 3, 9, 6},
			[]int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
		},
	}

	for _, tt := range tests {
		got := relativeSortArrayLessSpace(tt.arr1, tt.arr2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("failed for %#v: got=%#v\n", tt, got)
		}
	}
}

func TestRelativeSortArrayCS(t *testing.T) {
	tests := []TestCase{
		{
			[]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
			[]int{2, 1, 4, 3, 9, 6},
			[]int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
		},
	}

	for _, tt := range tests {
		got := relativeSortArrayCS(tt.arr1, tt.arr2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("failed for %#v: got=%#v\n", tt, got)
		}
	}
}
