package main

import "testing"

type TestCase struct {
	ratings []int
	want    int
}

func TestNumTeamsNaive(t *testing.T) {
	tests := []TestCase{
		{[]int{2, 5, 3, 4, 1}, 3},
		{[]int{2, 1, 3}, 0},
		{[]int{1, 2, 3, 4}, 4},
	}

	for _, tt := range tests {
		got := numTeamsNaive(tt.ratings)
		if got != tt.want {
			t.Errorf("failed for %#v: got=%d\n", tt, got)
		}
	}
}

func TestNumTeams(t *testing.T) {
	tests := []TestCase{
		{[]int{2, 5, 3, 4, 1}, 3},
		{[]int{2, 1, 3}, 0},
		{[]int{1, 2, 3, 4}, 4},
	}

	for _, tt := range tests {
		got := numTeams(tt.ratings)
		if got != tt.want {
			t.Errorf("failed for %#v: got=%d\n", tt, got)
		}
	}
}
