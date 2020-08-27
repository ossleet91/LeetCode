package main

// Find numbers with even number of digits.
//
// Given an array nums of integers, return how many of them contain an
// even number of digits.
//
// Example:
//
//	Input: nums = [12,345,2,6,7896]
//	Output: 2
//	Explanation:
//		12 contains 2 digits (even number of digits).
//		345 contains 3 digits (odd number of digits).
// 		2 contains 1 digit (odd number of digits).
// 		6 contains 1 digit (odd number of digits).
// 		7896 contains 4 digits (even number of digits).
//		Therefore only 12 and 7896 contain an even number of
//		digits.
//
// Constraints:
//
//	- 1 <= nums.length <= 500
//	- 1 <= nums[i] <= 10^5
//
// https://leetcode.com/problems/find-numbers-with-even-number-of-digits/

func countEvenDigits(nums []int) int {
	count := 0

	for _, n := range nums {
		// Problem description ensures maximum value in nums will be
		// 10^5 = 100,000.

		if (n >= 10 && n < 100) || (n >= 1_000 && n < 10_000) ||
			(n == 100_000) {
			count++
		}
	}

	return count
}

func main() {
}
