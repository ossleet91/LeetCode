package main

// Sort array by parity
//
// Given an array A of non-negative integers, return an array consisting
// of all the even elements of A, followed by all the odd elements of A.
//
// You may return any answer array that satisfies this condition.
//
// Example:
//	Input: [3,1,2,4]
//	Output: [2,4,3,1]
//	The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be
//	accepted.
//
// Constraints:
//
//	1. 1 <= A.length <= 5000
//	2. 0 <= A[i] <= 5000
//
// https://leetcode.com/problems/sort-array-by-parity/

import (
	"sort"
)

func paritySort(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i]%2 == 0
	})

	return arr
}

func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func paritySortTwoPtrs(arr []int) []int {
	for li, ri := 0, len(arr)-1; li < ri; {
		left, right := &arr[li], &arr[ri]

		if *left%2 != 0 {
			// Swap left with right if left is odd.
			//
			// Update the right index as we know the swapped
			// value is odd.
			//
			// Keep left index as is as the just swapped
			// value might be odd, for all we know.
			swap(left, right)
			ri--
		} else {
			// Current left value is even. So, move to next
			// left.
			li++
		}
	}

	return arr
}

// Sort given array by only parity (even/odd); even elements followed by
// odd.
//
// Uses Lomuto partition scheme algorithm.
// https://en.wikipedia.org/wiki/Quicksort#Lomuto_partition_scheme
//
// Time: O(N). Every element is looked almost once.
// Space: O(1). Swaps are done in place. Just two pointers are used.
func paritySortLomuto(arr []int) []int {
	for fast, slow := 0, 0; fast < len(arr); fast++ {
		if arr[fast]%2 == 0 {
			arr[fast], arr[slow] = arr[slow], arr[fast]
			slow++
		}
	}
	return arr
}

func main() {
}
