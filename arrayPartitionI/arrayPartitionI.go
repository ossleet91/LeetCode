package main

// 561. Array Partition I
//
// Given an array of 2n integers, your task is to group these integers
// into n pairs of integer, say (a1, b1), (a2, b2), ..., (an, bn) which
// makes sum of min(ai, bi) for all i from 1 to n as large as possible.
//
// Example 1:
//
//	Input: [1,4,3,2]
//	Output: 4
//	Explanation: n is 2, and the maximum sum of pairs is
//		     4 = min(1, 2) + min(3, 4).
//
// Note:
//
//     1. n is a positive integer, which is in the range of [1, 10000].
//     2. All the integers in the array will be in the range of
//	  [-10000, 10000].
//
// https://leetcode.com/problems/array-partition-i/description/

import (
	"math"
	"sort"
)

func minInArr(arr []int) int {
	min := math.MaxInt32
	for _, n := range arr {
		if n < min {
			min = n
		}
	}

	return min
}

func maxInArr(arr []int) int {
	max := math.MinInt32
	for _, n := range arr {
		if n > max {
			max = n
		}
	}

	return max
}

func countSort(arr []int) []int {
	count := make([]int, maxInArr(arr)+1)
	for _, n := range arr {
		count[n]++
	}

	total := 0
	for i, c := range count {
		count[i] = total
		total += c
	}

	out := make([]int, len(arr))
	for _, n := range arr {
		out[count[n]] = n
		count[n] += 1
	}

	return out
}

// Compute maximum sum of minimum of pairs from given array.
//
// i.e., if nums = []int{1, 3, 4, 2}, then, max. sum of min. of pairs is:
//	min(1,2) + min(3,4) = 1 + 3 = 4
//
// Any other pair-combination would not produce the maximum sum.
//
//	- (1,3), (4,2): 1 + 2 = 3
//	- (1,4), (3,2): 1 + 2 = 3
//
// Sort the array and then sum every alternate number from 0..n-1.
//
// Time: O(n + n.logn). n.logn to sort and n for summing.
//
// Space: O(1). Just running sum. stdlib's sort might use additioal
//        memory.
func arrayPairSumSort(nums []int) int {
	sort.Ints(nums)

	sum := 0
	for i := 0; i < len(nums); i = i + 2 {
		sum += nums[i]
	}

	return sum
}

// https://leetcode.com/problems/array-partition-i/discuss/182949/a-solution-by-go
//
// TODO: understand approach
func arrayPairSumMap(nums []int) int {
	countMap := make(map[int]int)
	for _, n := range nums {
		countMap[n]++
	}

	sum := 0
	for n, odd := -10000, false; n <= 10000; n, odd = n+1, !odd {
		for count := 0; count < countMap[n]; count++ {
			if odd {
				sum += n
			}
		}
	}

	return sum
}

func main() {
}
