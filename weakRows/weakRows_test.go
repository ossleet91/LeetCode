package main

import (
	"reflect"
	"testing"
)

func TestCountSoldiers(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{[]int{1, 1, 1, 1, 1}, 5},
		{[]int{1, 1, 1, 1, 0}, 4},
		{[]int{0, 0, 0, 0, 0, 0, 0}, 0},
	}

	for _, tt := range tests {
		got := countSoldiers(tt.arr)
		if got != tt.want {
			t.Errorf("failed for %+v: got=%d\n", tt, got)
		}
	}
}

func TestKWeakestRows(t *testing.T) {
	tests := []struct {
		mat  [][]int
		k    int
		want []int
	}{
		{
			[][]int{{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 0},
				{1, 0, 0, 0, 0},
				{1, 1, 0, 0, 0}},
			3,
			[]int{2, 0, 3},
		},
		{
			[][]int{{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 0},
				{1, 0, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 1}},
			3,
			[]int{2, 0, 3},
		},
	}

	for _, tt := range tests {
		got := kWeakestRows(tt.mat, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("failed for %+v: got=%+v\n", tt, got)
		}
	}
}
