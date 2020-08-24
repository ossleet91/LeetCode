package main

import (
	"reflect"
	"testing"
)

func TestCountOccurences(t *testing.T) {
	tests := []struct {
		nums []int
		val  int
		want int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 2, 3, 4, 5, 3, 1, 0}, 3, 2},
		{[]int{1, 1, 1}, 1, 3},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 3},
		{[]int{2}, 3, 0},
		{[]int{3, 3}, 3, 2},
	}

	for _, tt := range tests {
		got := countOccurences(tt.nums, tt.val)
		if tt.want != got {
			t.Errorf("failed for %+v: got=%d\n", tt, got)
		}
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		nums    []int
		val     int
		wantLen int
		want    []int
	}{
		{[]int{3, 2, 2, 3}, 3, 2, []int{2, 2}},
		{[]int{0, 2, 3, 4, 5, 3, 1, 0}, 3, 6, []int{0, 2, 0, 4, 5, 1}},
		{[]int{1, 1, 1}, 1, 0, []int{}},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5, []int{0, 1, 4, 0, 3}},
		{[]int{2}, 3, 1, []int{2}},
		{[]int{3, 3}, 5, 2, []int{3, 3}},
	}

	for _, tt := range tests {
		gotLen := removeElement(tt.nums, tt.val)
		if tt.wantLen != gotLen {
			t.Errorf("failed for %+v: gotLen=%d\n", tt, gotLen)
		}
		if !reflect.DeepEqual(tt.want, tt.nums[:tt.wantLen]) {
			t.Errorf("failed for %+v\n", tt)
		}
	}
}
