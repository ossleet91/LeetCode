package main

// Count negative numbers in a sorted matrix
//
// Given a m * n matrix grid which is sorted in non-increasing order
// both row-wise and column-wise.
//
// Return the number of negative numbers in grid.
//
// Example:
//
//	Input: grid = [[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]
//	Output: 8
//	Explanation: There are 8 negatives number in the matrix.
//
// https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/

import (
	"sort"
)

func countNegatives(grid [][]int) int {
	count := 0

	for i := 0; i < len(grid); i++ {
		row := grid[i]
		negIndex := sort.Search(len(row),
			func(index int) bool {
				return row[index] < 0
			})
		if negIndex < len(row) && row[negIndex] < 0 {
			count += len(row) - negIndex
		}
	}

	return count
}

func main() {
}
