package main

import (
	"testing"
)

type TestCase struct {
	n    int
	want int
}

func TestFibRecurse(t *testing.T) {
	tests := []TestCase{
		{2, 1},
		{3, 2},
		{4, 3},
		{27, 196418},
		{30, 832040},
	}

	for _, tt := range tests {
		got := fibRecurse(tt.n)
		if got != tt.want {
			t.Errorf("failed for %#v: got=%d\n", tt, got)
		}
	}
}

func TestFibLookup(t *testing.T) {
	tests := []TestCase{
		{2, 1},
		{3, 2},
		{4, 3},
		{27, 196418},
		{30, 832040},
	}

	for _, tt := range tests {
		got := fibLookup(tt.n)
		if got != tt.want {
			t.Errorf("failed for %#v: got=%d\n", tt, got)
		}
	}
}

func TestFib(t *testing.T) {
	tests := []TestCase{
		{2, 1},
		{3, 2},
		{4, 3},
		{27, 196418},
		{30, 832040},
	}

	for _, tt := range tests {
		got := fib(tt.n)
		if got != tt.want {
			t.Errorf("failed for %#v: got=%d\n", tt, got)
		}
	}
}
