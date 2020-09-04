package main

// Three consecutive odds
//
// Given an integer array arr, return true if there are three
// consecutive odd numbers in the array. Otherwise, return false.
//
// Example:
//
//	Input: arr = [2,6,4,1]
//	Output: false
//	Explanation: There are no three consecutive odds.
//
//	Input: arr = [1,2,34,3,4,5,7,23,12]
//	Output: true
//	Explanation: [5,7,23] are three consecutive odds.
//
// Constraints:
//
//	1. 1 <= arr.length <= 1000
//	2. 1 <= arr[i] <= 1000
//
// https://leetcode.com/problems/three-consecutive-odds

func threeConsecutiveOdds(arr []int) bool {
	for i := 0; i < len(arr)-2; i++ {
		if arr[i]%2 != 0 && arr[i+1]%2 != 0 && arr[i+2]%2 != 0 {
			return true
		}
	}

	return false
}

func main() {
}
