package main

import (
	"reflect"
	"testing"
)

func TestShuffleDup(t *testing.T) {
	tests := []struct {
		nums []int
		n    int
		want []int
	}{
		{[]int{2, 5, 1, 3, 4, 7}, 3, []int{2, 3, 5, 4, 1, 7}},
		{[]int{1, 2, 3, 4, 4, 3, 2, 1}, 4, []int{1, 4, 2, 3, 3, 2, 4, 1}},
		{[]int{1, 1, 2, 2}, 2, []int{1, 2, 1, 2}},
	}

	for _, tt := range tests {
		got := shuffleDup(tt.nums, tt.n)
		if !reflect.DeepEqual(tt.want, got) {
			t.Errorf("failed for %+v: got=%+v\n", tt.want, got)
		}
	}
}
