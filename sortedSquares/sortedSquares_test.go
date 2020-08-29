package main

import (
	"reflect"
	"testing"
)

type Test struct {
	arr  []int
	want []int
}

func TestSortedSquaresSort(t *testing.T) {
	tests := []Test{
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	}

	for _, tt := range tests {
		got := sortedSquaresSort(tt.arr)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("failed for %+v: got=%+v\n", tt, got)
		}
	}
}

func TestSortedSquares(t *testing.T) {
	tests := []Test{
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	}

	for _, tt := range tests {
		got := sortedSquares(tt.arr)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("failed for %+v: got=%+v\n", tt, got)
		}
	}
}
