package main

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestCountSort(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{3, 5, 7, 1, 8, 9}, []int{1, 3, 5, 7, 8, 9}},
	}

	for _, tt := range tests {
		got := countSort(tt.arr)
		if !reflect.DeepEqual(tt.want, got) {
			t.Errorf("failed for %+v: got=%+v\n", tt, got)
		}
	}
}

func TestCountRand(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length := rand.Intn(100000)
		arr, want := make([]int, length), make([]int, length)
		for i := range arr {
			arr[i] = rand.Intn(100000)
		}

		copy(want, arr)
		sort.Ints(want)

		got := countSort(arr)
		if !reflect.DeepEqual(want, got) {
			t.Errorf("failed for %+v: want=%+v got=%+v\n",
				arr, want, got)
		}

	}
}

type TestCase struct {
	nums []int
	want int
}

func TestArrayPairSumSort(t *testing.T) {
	tests := []TestCase{
		{[]int{1, 4, 3, 2}, 4},
	}

	for _, tt := range tests {
		nums := make([]int, len(tt.nums))
		copy(nums, tt.nums)

		got := arrayPairSumSort(nums)
		if got != tt.want {
			t.Errorf("failed for %+v: got=%d\n", tt, got)
		}
	}
}

func TestArrayPairSumMap(t *testing.T) {
	tests := []TestCase{
		{[]int{1, 4, 3, 2}, 4},
	}

	for _, tt := range tests {
		nums := make([]int, len(tt.nums))
		copy(nums, tt.nums)

		got := arrayPairSumMap(nums)
		if got != tt.want {
			t.Errorf("failed for %+v: got=%d\n", tt, got)
		}
	}
}
