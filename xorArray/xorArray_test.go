package main

import (
	"testing"
)

func TestXorOperation(t *testing.T) {
	tests := []struct {
		n     int
		start int
		want  int
	}{
		{5, 0, 8},
		{4, 3, 8},
		{1, 7, 7},
		{10, 5, 2},
	}

	for _, tt := range tests {
		got := xorOperation(tt.n, tt.start)
		if got != tt.want {
			t.Errorf("failed for %#v: got=%d\n", tt, got)
		}
	}
}
