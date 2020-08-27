package main

import (
	"testing"
)

func TestAbs(t *testing.T) {
	tests := []struct {
		val  int
		want int
	}{
		{1, 1},
		{-5000, 5000},
	}

	for _, tt := range tests {
		got := abs(tt.val)
		if got != tt.want {
			t.Errorf("failed for %+v: got=%d\n", tt, got)
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

func TestMinTimeToVisitAllPoints(t *testing.T) {
	tests := []struct {
		points [][]int
		want   int
	}{
		{[][]int{{1, 1}, {3, 4}, {-1, 0}}, 7},
		{[][]int{{3, 2}, {-2, 2}}, 5},
		{[][]int{{3, 2}}, 0},
	}

	for _, tt := range tests {
		got := minTimeToVisitAllPoints(tt.points)
		if got != tt.want {
			t.Errorf("failed for %+v: got=%d\n", tt, got)
		}
	}
}

func TestStep(t *testing.T) {
	tests := []struct {
		val  int
		dest int
		want int
	}{
		{3, 5, 4},
		{51, 30, 50},
		{5, 5, 5},
		{-55, 50, -54},
		{50, -55, 49},
	}

	for _, tt := range tests {
		got := step(tt.val, tt.dest)
		if got != tt.want {
			t.Errorf("failed for +%v: got=%d\n", tt, got)
		}
	}
}

func TestMinTimeToVisitAllPointsSlow(t *testing.T) {
	tests := []struct {
		points [][]int
		want   int
	}{
		{[][]int{{1, 1}, {3, 4}, {-1, 0}}, 7},
		{[][]int{{3, 2}, {-2, 2}}, 5},
		{[][]int{{3, 2}}, 0},
	}

	for _, tt := range tests {
		got := minTimeToVisitAllPointsSlow(tt.points)
		if got != tt.want {
			t.Errorf("failed for %+v: got=%d\n", tt, got)
		}
	}
}
