package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	heights []int
	want    int
}

func TestHeightCheckerSort(t *testing.T) {
	tests := []TestCase{
		{[]int{1, 1, 4, 2, 1, 3}, 3},
		{[]int{5, 1, 2, 3, 4}, 5},
		{[]int{1, 2, 3, 4, 5}, 0},
	}

	for _, tt := range tests {
		got := heightCheckerSort(tt.heights)
		if got != tt.want {
			t.Errorf("failed for %+v: got=%d", tt, got)
		}
	}
}

func TestMinMax(t *testing.T) {
	tests := []struct {
		nums    []int
		wantMin int
		wantMax int
	}{
		{[]int{}, 0, 0},
		{[]int{5}, 5, 5},
		{[]int{-10, 0, 10}, -10, 10},
		{[]int{10, 10, 10}, 10, 10},
	}

	for _, tt := range tests {
		gotMin, gotMax := minMax(tt.nums)
		if gotMin != tt.wantMin || gotMax != tt.wantMax {
			t.Errorf("failed for %+v: gotMin=%d gotMax=%d\n",
				tt, gotMin, gotMax)
		}
	}
}

func TestCountSort(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 1}, []int{1, 1}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{50, 40, 30, 20, 10}, []int{10, 20, 30, 40, 50}},
	}

	for _, tt := range tests {
		got := countSort(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("failed for %+v: got=%+v", tt, got)
		}
	}
}

func TestHeightChecker(t *testing.T) {
	tests := []TestCase{
		{[]int{1, 1, 4, 2, 1, 3}, 3},
		{[]int{5, 1, 2, 3, 4}, 5},
		{[]int{1, 2, 3, 4, 5}, 0},
	}

	for _, tt := range tests {
		got := heightChecker(tt.heights)
		if got != tt.want {
			t.Errorf("failed for %+v: got=%d", tt, got)
		}
	}
}
