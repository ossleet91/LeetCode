package main

import (
	"testing"
)

func TestNumJewelsInStones(t *testing.T) {
	tests := []struct {
		J    string
		S    string
		want int
	}{
		{"aA", "aAAbbbb", 3},
		{"z", "ZZ", 0},
	}

	for _, tt := range tests {
		got := numJewelsInStones(tt.J, tt.S)
		if tt.want != got {
			t.Errorf("failed for %+v: got=%d\n", tt, got)
		}
	}
}
