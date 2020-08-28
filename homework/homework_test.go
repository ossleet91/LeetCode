package main

import (
	"testing"
)

func TestBusyStudent(t *testing.T) {
	tests := []struct {
		startTime []int
		endTime   []int
		queryTime int
		want      int
	}{
		{[]int{1, 2, 3}, []int{3, 2, 7}, 4, 1},
		{[]int{4}, []int{4}, 4, 1},
		{[]int{4}, []int{4}, 5, 0},
		{[]int{1, 1, 1, 1}, []int{1, 3, 2, 4}, 7, 0},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{10, 10, 10, 10, 10, 10, 10, 10, 10}, 5, 5},
	}

	for _, tt := range tests {
		got := busyStudent(tt.startTime, tt.endTime, tt.queryTime)
		if got != tt.want {
			t.Errorf("failed for %+v: got=%d\n", tt, got)
		}
	}
}
