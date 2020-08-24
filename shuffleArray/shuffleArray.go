package main

// Shuffle the array
//
// Given the array nums consisting of 2n elements in the form
// [x1,x2,...,xn,y1,y2,...,yn].
//
// Return the array in the form [x1,y1,x2,y2,...,xn,yn].
//
// Example:
//
//	Input: nums = [2,5,1,3,4,7], n = 3
//	Output: [2,3,5,4,1,7]
//	Explanation: Since x1=2, x2=5, x3=1, y1=3, y2=4, y3=7 then the
//	answer is [2,3,5,4,1,7].
//
// https://leetcode.com/problems/shuffle-the-array/

func shuffleDup(nums []int, n int) []int {
	if len(nums) != 2*n {
		// If input length doesn't match problem's contract,
		// return the original array.
		return nums
	}

	res := make([]int, 2*n)
	for i, j := 0, 0; i < 2*n && j < 2*n; i, j = i+1, j+2 {
		res[j] = nums[i]
		res[j+1] = nums[i+n]
	}

	return res
}

func main() {
}
