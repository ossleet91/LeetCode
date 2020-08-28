package main

import (
	"reflect"
	"sort"
	"testing"
)

type Tests struct {
	target []int
	arr    []int
	want   bool
}

func TestHashSetDifference(t *testing.T) {
	tests := []struct {
		a    HashSet
		b    HashSet
		want []int
	}{
		{
			newValueHashSet([]int{1, 2, 2}),
			newValueHashSet([]int{2, 2, 1}),
			[]int{},
		},
		{
			newValueHashSet([]int{1, 2, 3, 4, 2}),
			newValueHashSet([]int{1, 2, 3, 3, 4}),
			[]int{2, 3},
		},
	}

	for _, tt := range tests {
		got := tt.a.Difference(tt.b)
		if len(tt.want) == 0 {
			if len(got) != 0 {
				t.Errorf("failed for empty %+v: got=%d\n",
					tt, got)
			}
		} else {
			// Set differnce can be in any order. So, sort,
			// and then compare.
			sort.Ints(tt.want)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("failed for %+v: got=%d\n", tt, got)
			}
		}
	}
}

func TestCanBeEqual(t *testing.T) {
	tests := []Tests{
		{[]int{1, 2, 3, 4}, []int{2, 4, 1, 3}, true},
		{[]int{7}, []int{7}, true},
		{[]int{1, 12}, []int{12, 1}, true},
		{[]int{3, 7, 9}, []int{3, 7, 11}, false},
		{[]int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1}, true},
	}

	for _, tt := range tests {
		got := canBeEqual(tt.target, tt.arr)
		if got != tt.want {
			t.Errorf("failed for %+v: got=%t\n", tt, got)
		}
	}
}
