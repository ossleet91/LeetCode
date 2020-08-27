package main

// XOR operation in an array.
//
// Given an integer n and an integer start.
//
// Define an array nums where nums[i] = start + 2*i (0-indexed) and n ==
// nums.length.
//
// Return the bitwise XOR of all elements of nums.
//
// Example:
//
//	Input: n = 5, start = 0
//	Output: 8
//	Explanation: Array nums is equal to [0, 2, 4, 6, 8] where
//		     (0 ^ 2 ^ 4 ^ 6 ^ 8) = 8. Where "^" corresponds to
//		     bitwise XOR operator.
//
// Constraints:
//
//	1 <= n <= 1000
//	0	<= start <= 1000
//	n == nums.length
//
// https://leetcode.com/problems/xor-operation-in-an-array/

func xorOperation(n, start int) int {
	res := 0

	for i := 0; i < n; i++ {
		res ^= start + (2 * i)
	}

	return res
}

func main() {
}
