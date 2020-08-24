package main

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	var tests = []struct {
		nums []int
		max  int
		want int
	}{
		{[]int{1, 1, 2, 3, 4, 5, 6, 6, 6, 7, 8, 9, 9}, 9, 9},
		{[]int{1, 1, 8}, 8, 2},
		{[]int{589, 589, 589}, 589, 1},
	}

	for _, tt := range tests {
		got := unique(tt.nums, tt.max)
		if got != tt.want {
			t.Errorf("%+v: got=%d\n", tt, got)
		}
	}
}

func TestRotateLeftTmp(t *testing.T) {
	var tests = []struct {
		array []int
		count int
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 0, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 2, 3, 4, 5, 6}, 1, []int{2, 3, 4, 5, 6, 1}},
		{[]int{1, 2, 3, 4, 5, 6}, 2, []int{3, 4, 5, 6, 1, 2}},
		{[]int{1, 2, 3, 4, 5, 6}, 3, []int{4, 5, 6, 1, 2, 3}},
		{[]int{1, 2, 3, 4, 5, 6}, 4, []int{5, 6, 1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5, 6}, 5, []int{6, 1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5, 6}, 6, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{4, 5, 6, 7, 1, 2, 3}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 7, []int{1, 2, 3, 4, 5, 6, 7}},
		{[]int{0, 5, 10, 15, 20, 25}, 0, []int{0, 5, 10, 15, 20, 25}},
		{[]int{0, 5, 10, 15, 20, 25}, 5, []int{25, 0, 5, 10, 15, 20}},
	}

	for _, tt := range tests {
		rotateLeftTmp(tt.array, tt.count)
		if !reflect.DeepEqual(tt.array, tt.want) {
			t.Errorf("failed for: %+v\n", tt)
		}

	}
}

func TestRotateLeftOne(t *testing.T) {
	var tests = []struct {
		array []int
		count int
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{4, 5, 6, 7, 1, 2, 3}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 7, []int{1, 2, 3, 4, 5, 6, 7}},
		{[]int{0, 5, 10, 15, 20, 25}, 5, []int{25, 0, 5, 10, 15, 20}},
		{[]int{0, 5, 10, 15, 20, 25}, 0, []int{0, 5, 10, 15, 20, 25}},
	}

	for _, tt := range tests {
		rotateLeftOne(tt.array, tt.count)
		if !reflect.DeepEqual(tt.array, tt.want) {
			t.Errorf("failed for: %+v\n", tt)
		}

	}
}

func TestRotateLeft(t *testing.T) {
	var tests = []struct {
		array []int
		count int
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{4, 5, 6, 7, 1, 2, 3}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 7, []int{1, 2, 3, 4, 5, 6, 7}},
		{[]int{0, 5, 10, 15, 20, 25}, 5, []int{25, 0, 5, 10, 15, 20}},
		{[]int{0, 5, 10, 15, 20, 25}, 0, []int{0, 5, 10, 15, 20, 25}},
	}

	for _, tt := range tests {
		rotateLeft(tt.array, tt.count)
		if !reflect.DeepEqual(tt.array, tt.want) {
			t.Errorf("failed for: %+v\n", tt)
		}

	}
}

func TestGcd(t *testing.T) {
	var tests = []struct {
		a    int
		b    int
		want int
	}{
		{7, 2, 1},
		{12, 3, 3},
	}

	for _, tt := range tests {
		got := gcd(tt.a, tt.b)
		if tt.want != got {
			t.Errorf("failed for: %+v: got=%d\n", tt, got)
		}

	}
}

func TestRemoveDuplicates1(t *testing.T) {
	var tests = []struct {
		nums    []int
		wantLen int
		want    []int
	}{
		{[]int{1}, 1, []int{1}},
		{[]int{1, 1, 1}, 1, []int{1}},
		{[]int{1, 5, 5}, 2, []int{1, 5}},
		{[]int{1, 1, 2}, 2, []int{1, 2}},
		{[]int{1, 1, 1, 3, 5, 6, 6, 8}, 5, []int{1, 3, 5, 6, 8}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
		{[]int{1, 5, 5, 5, 5, 5, 5}, 2, []int{1, 5}},
	}

	for _, tt := range tests {
		gotLen := removeDuplicates1(tt.nums)
		if gotLen != tt.wantLen {
			t.Errorf("failed for %+v: gotLen=%d\n", tt, gotLen)
		}

		if !reflect.DeepEqual(tt.want, tt.nums[:tt.wantLen]) {
			t.Errorf("arr=%+v wantLen=%d gotLen=%d got=%+v\n",
				tt.nums, tt.wantLen, gotLen, tt.nums[:tt.wantLen])
			t.Errorf("want=%+v got=%+v\n", tt.want, tt.nums[:tt.wantLen])
		}
	}
}
func TestRemoveDuplicates(t *testing.T) {
	var tests = []struct {
		nums    []int
		wantLen int
		want    []int
	}{
		{[]int{1}, 1, []int{1}},
		{[]int{1, 1, 1}, 1, []int{1}},
		{[]int{1, 5, 5}, 2, []int{1, 5}},
		{[]int{1, 1, 2}, 2, []int{1, 2}},
		{[]int{1, 1, 1, 3, 5, 6, 6, 8}, 5, []int{1, 3, 5, 6, 8}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
		{[]int{1, 5, 5, 5, 5, 5, 5}, 2, []int{1, 5}},
	}

	for _, tt := range tests {
		gotLen := removeDuplicates(tt.nums)
		if gotLen != tt.wantLen {
			t.Errorf("failed for %+v: gotLen=%d\n", tt, gotLen)
		}

		if !reflect.DeepEqual(tt.want, tt.nums[:tt.wantLen]) {
			t.Errorf("arr=%+v wantLen=%d gotLen=%d got=%+v\n",
				tt.nums, tt.wantLen, gotLen, tt.nums[:tt.wantLen])
			t.Errorf("want=%+v got=%+v\n", tt.want, tt.nums[:tt.wantLen])
		}
	}
}
