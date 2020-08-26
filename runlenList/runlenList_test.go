package main

import (
	"reflect"
	"testing"
)

func TestDecompressRLEList(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{2, 4, 4, 4}},
		{[]int{1, 1, 2, 3}, []int{1, 3, 3}},
		{[]int{4, 4, 4, 4}, []int{4, 4, 4, 4, 4, 4, 4, 4}},
	}

	for _, tt := range tests {
		got := decompressRLEList(tt.nums)
		if !reflect.DeepEqual(tt.want, got) {
			t.Errorf("failed for %#v: got=%#v\n", tt, got)
		}
	}
}
