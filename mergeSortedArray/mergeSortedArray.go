package main

// 88. Merge Sorted Array
//
// Given two sorted integer arrays nums1 and nums2, merge nums2 into
// nums1 as one sorted array.
//
// Note:
//
//     - The number of elements initialized in nums1 and nums2 are m and
//       n respectively.
//     - You may assume that nums1 has enough space (size that is equal
//       to m + n) to hold additional elements from nums2.
//
// Example:
//
//	Input:
//		nums1 = [1,2,3,0,0,0], m = 3
//		nums2 = [2,5,6],       n = 3
//
// 	Output: [1,2,2,3,5,6]
//
// Constraints:
//
//	1. -10^9 <= nums1[i], nums2[i] <= 10^9
//	2. nums1.length == m + n
//	3. nums2.length == n
//
// https://leetcode.com/problems/merge-sorted-array

import (
	"sort"
)

// Merge two sorted arrays.
//
// Given two sorted arrays, nums1 and nums2, and their lengths, m and n,
// merge nums2 into nums1 that resulting nums1 array remains sorted.
// nums1 has additional space to encompass nums2; i.e., len(nums1) =
// m+n.  Additional elements are all 0.
//
// Use a two-pointers approach: Start at the end of both arrays using
// the given length, and place the largest number at the end of nums1.
// Repeat this until one of the arrays is exhausted. If nums1 length is
// not-zero at the end of the loop, then, it is guaranteed that they are
// all smaller than nums2, and thus already in-place. If nums2 length is
// non-zero at the end of the loop, then, all remaining elements are
// smaller than what we already have/merged in nums1. So, just copy
// nums2[0..remainingLen] to num1[0..nums2_remainingLen].
//
// Time: O(m+n). All elements in nums1 and nums2 are visited at most
//	 once.
//
// Space: O(1). No additional array is allocated.
func mergeTwoPtrs(nums1 []int, m int, nums2 []int, n int) {
	// Nothing to merge if 2nd array is empty.
	if n == 0 {
		return
	}

	// Nothing to sort if 1st array length is 0. Just copy the 2nd
	// into first.
	if m == 0 {
		copy(nums1, nums2)
	}

	// Start both arrays from the end and place them from the end in
	// nums1.
	for m > 0 && n > 0 {
		if nums1[m-1] >= nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}

	// Copy any remaining element in nuns2. Order is
	// guaranteed as the the above loop took care of all larger
	// elements.
	if n > 0 {
		copy(nums1[:n], nums2[:n])
	}
}

// Merge two sorted arrays.
//
// Given two sorted arrays, nums1 and nums2, and their lengths, m and n,
// merge nums2 into nums1 that resulting nums1 array remains sorted.
// nums1 has additional space to encompass nums2; i.e., len(nums1) =
// m+n.  Additional elements are all 0.
//
// Use a naive approach: copy nums2 to nums1 at the end (i.e., from
// nums1[m..len(nums1)] and then sort the resulting array.
//
// Time: O(N.logN). Since array sort is used. stdlib might end up doing
//	 the sort in O(N).
//
// Space: O(1). No additional array is allocated.
func mergeNaive(nums1 []int, m int, nums2 []int, n int) {
	// Nothing to merge if 2nd array is empty.
	if n == 0 {
		return
	}

	// Nothing to sort if 1st array length is 0. Just copy the 2nd
	// into first.
	if m == 0 {
		copy(nums1, nums2)
	}

	// Copy nums2 at the end of nums1.
	//
	// nums1's actual length is m+n as constraint #2 guarantees
	// nums1 has enough space to hold all of nums2.
	copy(nums1[m:], nums2)

	// Sort the merged array to compute the result.
	sort.Ints(nums1)
}

func main() {
}
