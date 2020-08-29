package main

import (
	"reflect"
	"testing"
)

func TestMinInArr(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{[]int{1, 10, 4, 2}, 1},
		{[]int{9, 3, 8, 7}, 3},
		{[]int{15, -16, -17, 12}, -17},
		{[]int{4, 0, 17}, 0},
	}

	for _, tt := range tests {
		got := minInArr(tt.arr)
		if got != tt.want {
			t.Errorf("failed for %+v: got=%+v\n", tt, got)
		}
	}
}

func TestMaxInArr(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{[]int{1, 10, 4, 2}, 10},
		{[]int{9, 3, 8, 7}, 9},
		{[]int{15, -16, -17, 12}, 15},
		{[]int{4, 0, 17}, 17},
	}

	for _, tt := range tests {
		got := maxInArr(tt.arr)
		if got != tt.want {
			t.Errorf("failed for %+v: got=%+v\n", tt, got)
		}
	}
}

func TestGetColumn(t *testing.T) {
	tests := []struct {
		matrix   [][]int
		colIndex int
		want     []int
	}{
		{[][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}, 0, []int{3, 9, 15}},
		{[][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}, 1, []int{7, 11, 16}},
		{[][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}, 2, []int{8, 13, 17}},
		{[][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}, 0, []int{1, 9, 15}},
		{[][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}, 1, []int{10, 3, 16}},
		{[][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}, 2, []int{4, 8, 17}},
		{[][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}, 3, []int{2, 7, 12}},
	}

	for _, tt := range tests {
		got := getColumn(tt.matrix, tt.colIndex)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("failed for %+v: got=%+v\n", tt, got)
		}
	}
}

func TestLuckyNumbers(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   []int
	}{
		{[][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}, []int{15}},
		{[][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}, []int{12}},
	}

	for _, tt := range tests {
		got := luckyNumbers(tt.matrix)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("failed for %+v: got=%+v\n", tt, got)
		}
	}
}
