package main

import (
	"strings"
	"testing"
)

func TestDefangIPAddr(t *testing.T) {
	tests := []struct {
		addr string
		want string
	}{
		{"1.1.1.1", "1[.]1[.]1[.]1"},
		{"255.100.50.0", "255[.]100[.]50[.]0"},
	}

	for _, tt := range tests {
		got := defangIPAddr(tt.addr)
		if strings.Compare(tt.want, got) != 0 {
			t.Errorf("failed for %+v: got=%s\n", tt, got)
		}
	}
}
