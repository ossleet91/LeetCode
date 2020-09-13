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

func main() {
}
