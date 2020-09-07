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

func main() {
}
