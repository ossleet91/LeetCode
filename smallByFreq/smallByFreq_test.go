package main

import (
	"reflect"
	"testing"
)

func TestSmallCharCount(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"really", 1}, // 'a' is smallest
		{"seems", 2},  // 'e' is smallest
		{"cbd", 1},    // 'b' is smallest
		{"zaaaz", 3},  // 'a' is smallest
	}

	for _, tt := range tests {
		got := smallCharCount(tt.s)
		if got != tt.want {
			t.Errorf("failed for %+v: got=%d\n", tt, got)
		}
	}
}

type TestCase struct {
	queries []string
	words   []string
	want    []int
}

func TestNumSmallerByFrequencyNaive(t *testing.T) {
	tests := []TestCase{
		{
			[]string{"cbd"},
			[]string{"zaaaz"},
			[]int{1},
		},
		{
			[]string{"bbb", "cc"},
			[]string{"a", "aa", "aaa", "aaaa"},
			[]int{1, 2},
		},
	}

	for _, tt := range tests {
		got := numSmallerByFrequencyNaive(tt.queries, tt.words)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("failed for %#v: got=%#v\n", tt, got)
		}
	}
}

func TestNumSmallerByFrequency(t *testing.T) {
	tests := []TestCase{
		{
			[]string{"cbd"},
			[]string{"zaaaz"},
			[]int{1},
		},
		{
			[]string{"bbb", "cc"},
			[]string{"a", "aa", "aaa", "aaaa"},
			[]int{1, 2},
		},
	}

	for _, tt := range tests {
		got := numSmallerByFrequency(tt.queries, tt.words)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("failed for %#v: got=%#v\n", tt, got)
		}
	}
}
