package main

import (
	"reflect"
	"testing"
)

func TestFlipAndInvertImageSlow(t *testing.T) {
	tests := []struct {
		image [][]int
		want  [][]int
	}{
		{
			[][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}},
			[][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}},
		},
		{
			[][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}},
			[][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}},
		},
	}

	for _, tt := range tests {
		got := flipAndInvertImageSlow(tt.image)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("failed for %+v: got=%+v\n", tt, got)
		}
	}
}
