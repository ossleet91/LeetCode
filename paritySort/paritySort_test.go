package main

import (
	"testing"
)

type ParitySort struct {
	arr      []int // array to sort
	oddIndex int   // index at which odd begins; len(arr) represents no odd
}

func isAllEven(arr []int) bool {
	for _, n := range arr {
		if n%2 != 0 {
			return false
		}
	}

	return true
}

func isAllOdd(arr []int) bool {
	for _, n := range arr {
		if n%2 == 0 {
			return false
		}
	}

	return true
}

func TestParitySort(t *testing.T) {
	tests := []ParitySort{
		{[]int{3, 1, 2, 4}, 2},
		{[]int{55, 45, 36, 28, 21, 20, 15, 6, 3, 2}, 5},
		{[]int{0, 2, 4, 6, 8}, 5},
		{[]int{1, 3, 5}, 0},
	}

	for _, tt := range tests {
		got := paritySort(tt.arr)
		if !isAllEven(got[:tt.oddIndex]) {
			t.Errorf("odd mixed with even for %+v: got=%+v\n",
				tt, got)
		}
		if !isAllOdd(got[tt.oddIndex:]) {
			t.Errorf("even mixed with odd for %+v: got=%+v\n",
				tt, got)
		}
	}
}

func TestParitySortTwoPtrs(t *testing.T) {
	tests := []ParitySort{
		{[]int{3, 1, 2, 4}, 2},
		{[]int{55, 45, 36, 28, 21, 20, 15, 6, 3, 2}, 5},
		{[]int{0, 2, 4, 6, 8}, 5},
		{[]int{1, 3, 5}, 0},
	}

	for _, tt := range tests {
		got := paritySortTwoPtrs(tt.arr)
		if !isAllEven(got[:tt.oddIndex]) {
			t.Errorf("odd mixed with even for %+v: got=%+v\n",
				tt, got)
		}
		if !isAllOdd(got[tt.oddIndex:]) {
			t.Errorf("even mixed with odd for %+v: got=%+v\n",
				tt, got)
		}
	}
}

type CountMap map[int]int

func newCountMap(arr []int) CountMap {
	cm := CountMap{}
	for _, n := range arr {
		cm.Add(n)
	}
	return cm
}

func (cm CountMap) Add(n int) {
	cm[n]++
}

func (cm CountMap) Del(n int) {
	if val, ok := cm.Count(n); ok {
		if val > 1 {
			cm[n]--
		} else {
			cm.DelFull(n)
		}
	}
}

func (cm CountMap) Count(n int) (int, bool) {
	count, ok := cm[n]
	return count, ok
}

func (cm CountMap) DelFull(n int) {
	delete(cm, n)
}

func (cm CountMap) Present(n int) bool {
	_, ok := cm[n]
	return ok
}

func TestParitySortLomuto(t *testing.T) {
	tests := []ParitySort{
		{[]int{3, 1, 2, 4}, 2},
		{[]int{55, 45, 36, 28, 21, 20, 15, 6, 3, 2}, 5},
		{[]int{0, 2, 4, 6, 8}, 5},
		{[]int{1, 3, 5}, 0},
	}

	for _, tt := range tests {
		// Copy input as paritySortLomuto() does an in-place
		// sort.
		nums := make([]int, len(tt.arr))
		copy(nums, tt.arr)

		got := paritySortLomuto(nums)
		if !isAllEven(got[:tt.oddIndex]) {
			t.Errorf("odd mixed with even for %+v: got=%+v\n",
				tt, got)
		}
		if !isAllOdd(got[tt.oddIndex:]) {
			t.Errorf("even mixed with odd for %+v: got=%+v\n",
				tt, got)
		}

		// Ensure all numbers in input are present in output
		// with exact counts.
		wantCM := newCountMap(tt.arr)
		gotCM := newCountMap(got)
		for _, n := range tt.arr {
			wantCount, wantOK := wantCM.Count(n)
			gotCount, gotOK := gotCM.Count(n)

			if wantOK != gotOK {
				t.Errorf("%d presnet in want, but not in got; want=%+v got=%+v\n",
					n, tt.arr, got)
			}

			if wantCount != gotCount {
				t.Errorf("%d appears %d times in want, but %d in got; want=%+v got=%+v\n",
					n, wantCount, gotCount, tt.arr, got)
			}
		}
	}
}
