package main

import (
	"testing"
)

func TestAverageSalary(t *testing.T) {
	tests := []struct {
		salary []int
		want   float64
	}{
		{[]int{4000, 3000, 1000, 2000}, 2500.0},
		{[]int{1000, 2000, 3000}, 2000.0},
		{[]int{6000, 5000, 4000, 3000, 2000, 1000}, 3500.0},
		{[]int{8000, 9000, 2000, 3000, 6000, 1000}, 4750.0},
	}

	for _, tt := range tests {
		got := averageSalary(tt.salary)
		if got != tt.want {
			t.Errorf("failed for %+v: got=%f\n", tt, got)
		}
	}
}
