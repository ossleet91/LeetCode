package main

// Maximum product of two elements in an array.
//
// Given the array of integers nums, you will choose two different
// indices i and j of that array. Return the maximum value of
// (nums[i]-1)*(nums[j]-1).
//
// Example:
//
//	Input: nums = [3,4,5,2]
//	Output: 12
//	Explanation: If you choose the indices i=1 and j=2 (indexed from
//		     0), you will get the maximum value, that is,
//		     (nums[1]-1)*(nums[2]-1) = (4-1)*(5-1) = 3*4 = 12.
//
// https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/

func maxIndex(nums []int) (int, int) {
	max, index := 0, 0

	for i, n := range nums {
		if n > max {
			index = i
			max = n
		}
	}

	return index, max
}

func maxProduct(nums []int) int {
	// Find the first max value and then reset it so that we would
	// get the next max.
	i, ival := maxIndex(nums)
	nums[i] = 0

	_, jval := maxIndex(nums)

	return (ival - 1) * (jval - 1)
}

func main() {
}
