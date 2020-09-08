package main

// 1089. Duplicate Zeros
//
// Given a fixed length array arr of integers, duplicate each occurrence
// of zero, shifting the remaining elements to the right.
//
// Note that elements beyond the length of the original array are not
// written.
//
// Do the above modifications to the input array in place, do not return
// anything from your function.
//
// Example 1:
//
//	Input: [1,0,2,3,0,4,5,0]
// 	Output: null
// 	Explanation: After calling your function, the input array is
//		     modified to: [1,0,0,2,3,0,0,4]
//
// Example 2:
//
//	Input: [1,2,3]
// 	Output: null
// 	Explanation: After calling your function, the input array is
// 	             modified to: [1,2,3]
//
// Note:
//
//	1. 1 <= arr.length <= 10000
// 	2. 0 <= arr[i] <= 9
//
// https://leetcode.com/problems/duplicate-zeros

// Right shift a given array whose length is at least 2 by 1 postion.
// The first element is left untouched. The last element is left-out.
//
// For e.g., arr=[]int{1,2,3,4} after shift becomes []int{1,1,2,3}.
//
// Time: O(N). All elements are processed at most once.
//
// Space: O(1). Shfiting is done in-place.
func shift(arr []int) {
	if len(arr) == 1 {
		return
	}

	for i := len(arr) - 1; i >= 1; i-- {
		arr[i] = arr[i-1]
	}
}

// Duplicate zeros in an array in-place.
//
// For every zero at index i, shfit remaining elements to i+2..n and add
// another zero at i+1. Any elements beyond the length of the array are
// left-out.
//
// For e.g., a=[]int{1,2,0,3,0,4} becomes []int{1,2,0,0,3,0}.
//
// Time: O(N^2). Every element is scanned once to check if it's a zero
//       and if zero, remaining elements are shifted using O(N)
//       approach.
//
// Space: O(1). All shifts/inserts are done in-place.
func duplicateZerosNaive(arr []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 && i < len(arr)-1 {
			shift(arr[i:])
			i++ // increment i to account for new zero
		}
	}
}

// Duplicate zeros in an array in-place.
//
// Use two pointers: origIndex, to keep track of element in the given
// array, and updatedIndex, to track where the element should actually
// go (this accounts for zero duplications).
//
// Based on:
// https://leetcode.com/problems/duplicate-zeros/discuss/315395/Java-O(n)-Time-O(1)-space-Two-passes-for-loop-and-while-loop
//
// Time: O(N). Array is looped linearly once.
//
// Space: O(1). Only couple of pointers and a counter.
func duplicateZerosTwoPtrs(arr []int) {
	// Count all zeros. Used to compute the final position of each
	// element once we start duplicating the elements.
	zeroCount := 0
	for _, n := range arr {
		if n == 0 {
			zeroCount++
		}
	}

	// origIndex is the original index of every element in the
	// array; for 0..len(arr)-1.
	origIndex := len(arr) - 1

	// updatedIndex tracks the index at which a given number is
	// supposed to go at.
	//
	// Since it accounts for the additional zeros that we have to
	// duplicate, the value might be outside the array bounds. So,
	// check for bounds every time when this index is used to set an
	// element.
	updatedIndex := len(arr) + zeroCount - 1

	for origIndex >= 0 && updatedIndex >= 0 {
		n := arr[origIndex]

		// Insert the current element at the appropriate index.
		if updatedIndex < len(arr) {
			arr[updatedIndex] = n
		}

		if n == 0 {
			// Current element is a zero and we have to
			// duplicate if index is within array-bound.
			//
			// So, update current position and duplicate the
			// zero. Again, check position is within array
			// bounds.
			updatedIndex--
			if updatedIndex < len(arr) {
				arr[updatedIndex] = n
			}
		}

		// Go to next element. This could have been done in the
		// 'for', but since init was moved out (due to verbose
		// comments), moved post section out of the 'for' as
		// well.
		origIndex--
		updatedIndex--
	}
}

func main() {
}
