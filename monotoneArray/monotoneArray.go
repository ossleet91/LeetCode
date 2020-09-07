package main

// 896. Monotonic Array
//
//
// An array is monotonic if it is either monotone increasing or monotone
// decreasing.
//
// An array A is monotone increasing if for all i <= j, A[i] <= A[j].
// An array A is monotone decreasing if for all i <= j, A[i] >= A[j].
//
// Return true if and only if the given array A is monotonic.
//
//
// Example 1:
//
//	Input: [1,2,2,3]
// 	Output: true
//
// Example 2:
//
// 	Input: [6,5,4,4]
// 	Output: true
//
// Example 3:
//
// 	Input: [1,3,2]
// 	Output: false
//
// Example 4:
//
// 	Input: [1,2,4,5]
// 	Output: true
//
// Example 5:
//
// 	Input: [1,1,1]
// 	Output: true
//
//
// Note:
//
//	1. 1 <= A.length <= 50000
//	2. -100000 <= A[i] <= 100000
//
//
// https://leetcode.com/problems/monotonic-array

// Check whether given arr is monotonic or not.
//
// Record increase/decrease for the first time when arr[i] and arr[i-1]
// differ. Ensure this constraint holds for every other number that's
// different from it's left-negihbour.
//
// Time: O(N). We look at every number in the array.
//
// Space: O(1). Only one local variable to record first
//        increase/decrease.
func isMonotonic(arr []int) bool {
	if len(arr) == 1 {
		return true
	}

	change := 0
	for i := 1; i < len(arr); i++ {
		if arr[i-1] == arr[i] {
			continue
		}

		if change == 0 {
			if arr[i-1] < arr[i] {
				change = -1 // decreasing
			} else {
				change = 1 // increasing
			}
		}

		if change == -1 && arr[i-1] > arr[i] {
			// Decrease followed by an increase. Not
			// monotonic.
			return false
		}

		if change == 1 && arr[i-1] < arr[i] {
			// Increase followed by a decrease. Not
			// monotonic.
			return false
		}
	}

	return true
}

func main() {
}
