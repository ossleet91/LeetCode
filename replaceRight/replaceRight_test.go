package main

import (
	"reflect"
	"testing"
)

func TestMaxIndex(t *testing.T) {
	tests := []struct {
		nums      []int
		wantIndex int
		wantMax   int
	}{
		{[]int{3, 4, 5, 2}, 2, 5},
		{[]int{1, 5, 4, 5}, 1, 5},
		{[]int{3, 7}, 1, 7},
	}

	for _, tt := range tests {
		gotIndex, gotMax := maxIndex(tt.nums)
		if gotIndex != tt.wantIndex || gotMax != tt.wantMax {
			t.Errorf("failed for %+v: gotIndex=%d gotMax=%d\n",
				tt, gotIndex, gotMax)
		}
	}
}

func TestReplaceRightSlow(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{17, 18, 5, 4, 6, 1}, []int{18, 6, 6, 6, 1, -1}},
	}

	for _, tt := range tests {
		got := replaceRightSlow(tt.arr)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("failed for %+v: got=%+v\n", tt, got)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{1, 1, 1},
		{-5000, 5000, 5000},
		{-5000, -5001, -5000},
		{9, 81, 81},
	}

	for _, tt := range tests {
		got := max(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("failed for %+v: got=%d\n", tt, got)
		}
	}
}

func TestReplaceRight(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{17, 18, 5, 4, 6, 1}, []int{18, 6, 6, 6, 1, -1}},
	}

	for _, tt := range tests {
		got := replaceRight(tt.arr)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("failed for %+v: got=%+v\n", tt, got)
		}
	}
}
