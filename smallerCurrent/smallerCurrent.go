package main

// How many numbers are smaller than the current number
//
// Given the array nums, for each nums[i] find out how many numbers in
// the array are smaller than it. That is, for each nums[i] you have to
// count the number of valid j's such that j != i and nums[j] < nums[i].
//
// Return the answer in an array.
//
// Example:
//	Input: nums = [8,1,2,2,3]
//	Output: [4,0,1,1,3]
//
// Explanation:
//	For nums[0]=8 there exist four smaller numbers than it (1, 2, 2 and 3).
//	For nums[1]=1 does not exist any smaller number than it.
//	For nums[2]=2 there exist one smaller number than it (1).
//	For nums[3]=2 there exist one smaller number than it (1).
//	For nums[4]=3 there exist three smaller numbers than it (1, 2 and 2).
//
// https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/

// For every element, nums[i], check every other element, nums[j], and
// update a count, small[i] if nums[j] < nums[i].
//
// Runtime: O(n^2)
// Space: O(n)
func smallerThanCurrent(nums []int) []int {
	small := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			}

			if nums[j] < nums[i] {
				small[i]++
			}
		}
	}

	return small
}

// Since the values in nums are in range [0..100], we can use counting
// sort to find the smaller numbers.
//
// Runtime: O(n)
// Space: O(1)
func smallerThanCurrentCountSort(nums []int) []int {
	// Record the frequency of each element in 'nums'.
	//
	// The problem constraint guarantees the elemenets in 'nums' to
	// be in the range [0..100]. So, allocate memory for 101 ints.
	freqCount := make([]int, 101)
	for _, n := range nums {
		freqCount[n]++
	}

	// Replace the frequency of each element with a running counter.
	count := 0
	for i, f := range freqCount {
		if f != 0 {
			freqCount[i] = count
			count += f
		}
	}

	// Replace original input with freqCount value.
	for i := 0; i < len(nums); i++ {
		nums[i] = freqCount[nums[i]]
	}

	return nums
}

func main() {
}
