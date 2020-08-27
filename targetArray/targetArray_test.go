package main

import (
	"reflect"
	"testing"
)

func TestCreateTargetArrary(t *testing.T) {
	tests := []struct {
		nums  []int
		index []int
		want  []int
	}{
		{[]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}, []int{0, 4, 1, 3, 2}},
		{[]int{1, 2, 3, 4, 0}, []int{0, 1, 2, 3, 0}, []int{0, 1, 2, 3, 4}},
		{[]int{4, 2, 4, 3, 2}, []int{0, 0, 1, 3, 1}, []int{2, 2, 4, 4, 3}},
	}

	for _, tt := range tests {
		got := createTargetArray(tt.nums, tt.index)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("failed for %#v: got=%#v\n", tt, got)
		}
	}
}
