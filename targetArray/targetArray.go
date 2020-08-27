package main

// Create target array in given order
//
// Given two arrays of integers nums and index. Your task is to create
// target array under the following rules:
//
//	- Initially target array is empty.
//	- From left to right read nums[i] and index[i], insert at index
//	  index[i] the value nums[i] in target array.
//	- Repeat the previous step until there are no elements to read
//	  in nums and index.
//
// Return the target array.
//
// It is guaranteed that the insertion operations will be valid.
//
// Example:
//
//	Input: nums = [0,1,2,3,4], index = [0,1,2,2,1]
//	Output: [0,4,1,3,2]
//	Explanation:
//		nums       index     target
//		0            0        [0]
//		1            1        [0,1]
// 		2            2        [0,1,2]
// 		3            2        [0,1,3,2]
// 		4            1        [0,4,1,3,2]
//
// Constraints:
//
//	- 1 <= nums.length, index.length <= 100
//	- nums.length == index.length
//	- 0 <= nums[i] <= 100
//	- 0 <= index[i] <= i
//
//
// https://leetcode.com/problems/create-target-array-in-the-given-order/

func createTargetArray(nums, index []int) []int {
	seen := make([]bool, len(nums))
	res := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		val, targetIndex := nums[i], index[i]
		if !seen[targetIndex] {
			// We have never seen targetIndex before. So,
			// mark it as seen and set the value in result
			// array.
			seen[targetIndex] = true
			res[targetIndex] = val
		} else {
			// For indices that we had seen before, there's
			// already a value at the target index. So, we
			// have to shift our array to the right to make
			// room for the current value.
			copy(res[targetIndex+1:], res[targetIndex:])
			res[targetIndex] = val

			// We just touched all the indices following the
			// current index. So, set their 'seen' flag to
			// true.
			for j := targetIndex + 1; j < len(seen); j++ {
				seen[j] = true
			}
		}
	}

	return res
}

func main() {
}
