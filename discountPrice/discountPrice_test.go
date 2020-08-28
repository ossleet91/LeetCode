package main

import (
	"reflect"
	"testing"
)

func TestDiscountPriceSlow(t *testing.T) {
	tests := []struct {
		prices []int
		want   []int
	}{
		{[]int{8, 4, 6, 2, 3}, []int{4, 2, 4, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{10, 1, 1, 6}, []int{9, 0, 1, 6}},
		{[]int{8, 7, 4, 2, 8, 1, 7, 7, 10, 1}, []int{1, 3, 2, 1, 7, 0, 0, 6, 9, 1}},
	}

	for _, tt := range tests {
		got := discountPriceSlow(tt.prices)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("failed for %+v: got=%+v\n", tt, got)
		}
	}
}

func TestDiscountPrice(t *testing.T) {
	tests := []struct {
		prices []int
		want   []int
	}{
		{[]int{8, 4, 6, 2, 3}, []int{4, 2, 4, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{10, 1, 1, 6}, []int{9, 0, 1, 6}},
		{[]int{8, 7, 4, 2, 8, 1, 7, 7, 10, 1}, []int{1, 3, 2, 1, 7, 0, 0, 6, 9, 1}},
	}

	for _, tt := range tests {
		got := discountPrice(tt.prices)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("failed for %+v: got=%+v\n", tt, got)
		}
	}
}
