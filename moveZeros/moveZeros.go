package main

// 283. Move Zeros
//
// Given an array nums, write a function to move all 0's to the end of
// it while maintaining the relative order of the non-zero elements.
//
// Example:
//
//	Input: [0,1,0,3,12]
//	Output: [1,3,12,0,0]
//
//
// Note:
//
//	1. You must do this in-place without making a copy of the array.
//	2. Minimize the total number of operations.
//
//
// https://leetcode.com/problems/move-zeroes

// Move all zeros in given array to the end whilst maintaining relative
// positions of non-zero elements.
//
// Use two-pointer, slow & fast, apporach. Slow and fast starts at 0,
// fast is used as array-iterator (i.e., incremented at every
// iteration), and slow is used to point the first zero.
//
// When fast points to non-zero, swap slow/fast elements.
// Increment slow only  pointer keeps
// track of first-zero (relative) to swap with fast.
//
// Time: O(N). All elements of the array are visited at most once.
//
// Space: O(1). Only slow/fast pointers are used.
func moveZeros(nums []int) {
	for slow, fast := 0, 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
		}

		// Increment slow in the current iteration if it's not
		// pointing to 0. If already 0, slow already points to
		// the right index for swapping. So, don't increment
		// (and hence the name slow).
		if nums[slow] != 0 {
			slow++
		}
	}
}

// Move all zeros in given array to the end whilst maintaining relative
// positions of non-zero elements.
//
// Uses Lomuto partition sheme for moving around elements.
// https://en.wikipedia.org/wiki/Quicksort#Lomuto_partition_scheme
//
// It is simple, but swaps the same non-zero number.
//
// Time: O(N). All elements are visited at most once.
//
// Space: O(1). Only slow/fast pointers are used.
func moveZerosLomuto(nums []int) {
	for slow, fast := 0, 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}
	}
}

func main() {
}
