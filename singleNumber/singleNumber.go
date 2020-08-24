package main

// Single number problem
//
// Given a non-empty array of integers, every element appears twice
// except for one. Find that single one.
//
// Note:
// Your algorithm should have a linear runtime complexity. Could you
// implement it without using extra memory?
//
// Example:
//	Input: []int{2,2,1}
//	Output: 1
//
// https://leetcode.com/problems/single-number/

func singleNumber(nums []int) int {
	val := 0
	for _, n := range nums {
		val ^= n
	}

	return val
}

func main() {
}
