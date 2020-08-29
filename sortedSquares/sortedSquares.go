package main

// Squares of a Sorted Array
//
// Given an array of integers A sorted in non-decreasing order, return
// an array of the squares of each number, also in sorted non-decreasing
// order.
//
// Example:
//	Input: [-4,-1,0,3,10]
//	Output: [0,1,9,16,100]
//
// Constraints:
//
//	1. 1 <= A.length <= 10000
//	2. -10000 <= A[i] <= 10000
//	3. A is sorted in non-decreasing order.
//
// https://leetcode.com/problems/squares-of-a-sorted-array/

import (
	"fmt"
	"sort"
)

func sortedSquaresSort(arr []int) []int {
	for i, n := range arr {
		arr[i] = n * n
	}

	sort.Ints(arr)

	return arr
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func sortedSquares(arr []int) []int {
	res := make([]int, len(arr))

	for i, j, k := 0, len(arr)-1, len(arr)-1; k >= 0; k-- {
		left, right := arr[i], arr[j]

		if abs(left) > abs(right) {
			res[k] = left * left
			i++
		} else {
			res[k] = right * right
			j--
		}
	}

	return res
}

func main() {
	fmt.Println("vim-go")
}
