package main

// Remove a given element in place.
//
// Given an array nums and a value val, remove all instances of that
// value in-place and return the new length.
//
// Do not allocate extra space for another array, you must do this by
// modifying the input array in-place with O(1) extra memory.
//
// The order of elements can be changed. It doesn't matter what you
// leave beyond the new length.
//
// Example:
//
// Given nums = [3,2,2,3], val = 3,
//
// Your function should return length = 2, with the first two elements
// of nums being 2.
//
// It doesn't matter what you leave beyond the returned length.
//
// https://leetcode.com/problems/remove-element/

func swap(nums []int, fwd, bck int) int {
	for nums[fwd] == nums[bck] && bck > fwd {
		bck--
	}

	tmp := nums[bck]
	nums[bck] = nums[fwd]
	nums[fwd] = tmp

	return bck
}

func countOccurences(nums []int, val int) int {
	count := 0

	for _, n := range nums {
		if n == val {
			count++
		}
	}

	return count
}

func removeElement(nums []int, val int) int {
	valCount := countOccurences(nums, val)
	if valCount == 0 {
		return len(nums)
	} else if valCount == len(nums) {
		return 0
	}

	bck := len(nums) - 1
	for fwd := 0; fwd < len(nums); fwd++ {
		if nums[fwd] == val {
			bck = swap(nums, fwd, bck)
		}
	}

	return bck
}

func main() {
}
