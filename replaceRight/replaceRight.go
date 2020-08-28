package main

// Rpelace elements with greatest element on right side.
//
// Given an array arr, replace every element in that array with the
// greatest element among the elements to its right, and replace the
// last element with -1.
//
// After doing so, return the array.
//
// Example:
//
//	Input: arr = [17,18,5,4,6,1]
//	Output: [18,6,6,6,1,-1]
//
// Constraints:
//
//	- 1 <= arr.length <= 10^4
//	- 1 <= arr[i] <= 10^5
//
// https://leetcode.com/problems/replace-elements-with-greatest-element-on-right-side/

import "fmt"

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

func replaceRightSlow(arr []int) []int {
	index, max := 0, 0

	for i := range arr {
		if i == len(arr)-1 {
			arr[i] = -1
			break
		}

		if i == 0 || i > index {
			index, max = maxIndex(arr[i+1:])
		}
		arr[i] = max
	}

	return arr
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func replaceRight(arr []int) []int {
	maxVal := -1

	for i := len(arr) - 1; i >= 0; i-- {
		tmp := arr[i]
		arr[i] = maxVal
		maxVal = max(maxVal, tmp)
	}

	return arr
}

func main() {
	fmt.Println("vim-go")
}
