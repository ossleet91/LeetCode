package main

import (
	"testing"
)

func TestCountNegatives(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}, 8},
		{[][]int{{3, 2}, {1, 0}}, 0},
		{[][]int{{1, -1}, {-1, -1}}, 3},
		{[][]int{{-1}}, 1},
	}

	for _, tt := range tests {
		got := countNegatives(tt.grid)
		if got != tt.want {
			t.Errorf("failed for %+v: got=%d\n", tt, got)
		}
	}
}
