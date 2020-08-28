package main

// Find n unique integers that sum to zero
//
// Given an integer n, return any array containing n unique integers
// such that they add up to 0.
//
// Example:
//	Input: n = 5
//	Output: [-7,-1,1,3,4]
//	Explanation: These arrays also are accepted [-5,-1,1,2,3] ,
//		     [-3,-1,2,-2,4].
//
// Constraints
//	- 1 <= n <= 1000
//
// https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero/

import "fmt"

func zeroSum(n int) []int {
	if n == 1 {
		return []int{0}
	}

	sum := 0
	res := make([]int, n)

	// Fill up i=[0..n-2] positions with 1..i+n-1 values.
	// Fill the last position with negative sum of all the previous
	// values.
	//
	// i.e., for n=5, return []int{1, 2, 3, 4, -10}.
	for i := 0; i < n-1; i++ {
		res[i] = i + 1
		sum += res[i]
	}
	res[n-1] = -sum

	return res
}

func main() {
	fmt.Println("vim-go")
}
