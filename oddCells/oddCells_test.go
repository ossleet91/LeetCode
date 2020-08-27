package main

import (
	"testing"
)

func TestCountOddCells(t *testing.T) {
	tests := []struct {
		n       int
		m       int
		indices [][]int
		want    int
	}{
		{2, 3, [][]int{{0, 1}, {1, 1}}, 6},
		{2, 2, [][]int{{1, 1}, {0, 0}}, 0},
	}

	for _, tt := range tests {
		got := countOddCells(tt.n, tt.m, tt.indices)
		if got != tt.want {
			t.Errorf("failed for %+v: got=%d\n", tt, got)
		}
	}
}
