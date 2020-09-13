package main

// 1051. Height Checker
//
// Students are asked to stand in non-decreasing order of heights for an
// annual photo.
//
// Return the minimum number of students that must move in order for all
// students to be standing in non-decreasing order of height.
//
// Notice that when a group of students is selected they can reorder in
// any possible way between themselves and the non selected students
// remain on their seats.
//
// Example 1:
//
//	Input: heights = [1,1,4,2,1,3]
// 	Output: 3
// 	Explanation:
// 	Current array : [1,1,4,2,1,3]
// 	Target array  : [1,1,1,2,3,4]
// 	On index 2 (0-based) we have 4 vs 1 so we have to move this student.
// 	On index 4 (0-based) we have 1 vs 3 so we have to move this student.
// 	On index 5 (0-based) we have 3 vs 4 so we have to move this student.
//
// Example 2:
//
//	Input: heights = [5,1,2,3,4]
// 	Output: 5
//
// Example 3:
//
//	Input: heights = [1,2,3,4,5]
// 	Output: 0
//
// Constraints:
//
//     1. 1 <= heights.length <= 100
//     2. 1 <= heights[i] <= 100
//
// https://leetcode.com/problems/height-checker

import (
	"math"
	"sort"
)

// Sort the given array and count no. of positions that differ between
// sorted and unsorted array. The count is the total no. of moves
// required to arrange students by height.
//
// Time: O(N * logN). Worst case time to sort the array.
//
// Space: O(N). Extra space for sorted array.
func heightCheckerSort(heights []int) int {
	sorted := make([]int, len(heights))
	copy(sorted, heights)
	sort.Ints(sorted)

	res := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != sorted[i] {
			res++
		}
	}

	return res
}

func minMax(nums []int) (int, int) {
	switch len(nums) {
	case 0:
		return 0, 0
	case 1:
		return nums[0], nums[0]
	default:
		min, max := math.MaxInt32, math.MinInt32

		for _, n := range nums {
			if n < min {
				min = n
			}

			if n > max {
				max = n
			}
		}

		return min, max
	}
}

func doCountSort(nums []int, max int) []int {
	if len(nums) < 2 {
		return nums
	}

	// Compute frequency of each element and its prefix sum.
	prefixSum := make([]int, max+1)
	for _, n := range nums {
		prefixSum[n]++
	}

	for i := 1; i < len(prefixSum); i++ {
		prefixSum[i] += prefixSum[i-1]
	}

	// Do counting sort.
	//
	// For each element in given array, n, find it's prefix sum. The
	// index in sorted array for 'n' is its prefixSum-1; i.e., if
	// prefixSum is 5, then 'n' should be placed at index 4.
	//
	// Decrement the prefixSum so that we can handle duplicates.
	sorted := make([]int, len(nums))
	for _, n := range nums {
		index := prefixSum[n] - 1
		sorted[index] = n
		prefixSum[n]--
	}

	return sorted
}

func countSort(nums []int) []int {
	_, max := minMax(nums)
	return doCountSort(nums, max)
}

// Sort the given array and count no. of positions that differ between
// sorted and unsorted array. The count is the total no. of moves
// required to arrange students by height.
//
// Time: O(N+k). Since counting sort is used for sorting.
//
// Space: O(N). Extra space for sorted array.
func heightChecker(heights []int) int {
	// 100 is the max. value in heights. See constraint #2.
	sorted := doCountSort(heights, 100)

	res := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != sorted[i] {
			res++
		}
	}

	return res
}

func main() {
}
