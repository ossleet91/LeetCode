package main

import (
	"reflect"
	"testing"
)

type TopKFreq struct {
	nums []int
	k    int
	want []int
}

func TestTopKFreq(t *testing.T) {
	tests := []TopKFreq{
		{[]int{1, 1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 2}, 2, []int{1, 2}},
		{[]int{1, 1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4}, 2, []int{1, 2}},
	}

	for _, tt := range tests {
		got := topKFreq(tt.nums, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("failed for %+v: got=%d\n", tt, got)
		}
	}
}

func TestTopKFreqPQ(t *testing.T) {
	tests := []TopKFreq{
		{[]int{1, 1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 2}, 2, []int{1, 2}},
		{[]int{1, 1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4}, 2, []int{1, 3}},
	}

	for _, tt := range tests {
		got := topKFreqPQ(tt.nums, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("failed for %+v: got=%d\n", tt, got)
		}
	}
}
