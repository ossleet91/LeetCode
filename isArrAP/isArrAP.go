package main

// Can make arithmetic progression from sequence?
//
// Given an array of numbers arr. A sequence of numbers is called an
// arithmetic progression if the difference between any two consecutive
// elements is the same.
//
// Return true if the array can be rearranged to form an arithmetic
// progression, otherwise, return false.
//
//
//
// Example:
//
//	Input: arr = [3,5,1]
//	Output: true
//	Explanation: We can reorder the elements as [1,3,5] or [5,3,1]
//		     with differences 2 and -2 res
//
// Constraints:
//
//	1. 2 <= arr.length <= 1000
//	2. -10^6 <= arr[i] <= 10^6
//
// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence/

import (
	"sort"
)

func canMakeAPSort(arr []int) bool {
	sort.Ints(arr)

	diff := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		curr, prev := arr[i], arr[i-1]
		if curr-prev != diff {
			return false
		}
	}

	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

type HashSet map[int]bool

func newHashSet() HashSet {
	return HashSet{}
}

func (hs HashSet) Add(n int) {
	hs[n] = true
}

func (hs HashSet) Contains(n int) bool {
	_, ok := hs[n]
	return ok
}

// AP: v, v+1.d, v+2.d, v+3.d, ..., v+(n-1).d
//
// Let T1, Tn denote the first and last terms in the AP.
// Thus, we can find n'th term using: Tn = T1 + (n-1)d
//
// Then, Tn - T1 = d(n-1) or d = (Tn - Tn-1) / (n - 1)	==> FORMULA
//
// So, we can use a set to check a given array can form AP or not in
// O(n) without sorting.
//
// 1. Add all elements to a set. Use this loop to find min/max elements.
// 2. Given min/max, we can find 'd'. See FORMULA above.
// 3. Given min, n, and d, find the next suppossed element is the AP.
// Check for it in set. If not present, this array cannot form an AP. If
// present, move to next term in AP until n terms. If all terms are
// present in set, then, the given array clearly forms an AP.
func canMakeAP(arr []int) bool {
	// Use min/max values in constraint #2 as init values.
	minv, maxv := 1_000_000, -1_000_000
	set := newHashSet()
	for _, n := range arr {
		minv = min(n, minv)
		maxv = max(n, maxv)
		set.Add(n)
	}

	delta := (maxv - minv) / (len(arr) - 1)
	for i := 0; i < len(arr); i++ {
		if !set.Contains(minv) {
			return false
		}

		// Next supposed term in an AP is always
		// curr_term + delta.
		minv += delta
	}

	return true
}

func main() {
}
