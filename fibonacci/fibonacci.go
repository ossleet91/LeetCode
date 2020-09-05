package main

// 509. Fibonacci number
//
// The Fibonacci numbers, commonly denoted F(n) form a sequence, called the
// Fibonacci sequence, such that each number is the sum of the two preceding
// ones, starting from 0 and 1. That is,
//
//	F(0) = 0,   F(1) = 1
//	F(N) = F(N - 1) + F(N - 2), for N > 1.
//
// Given N, calculate F(N).
//
// Example 1:
//
//	Input: 2
// 	Output: 1
// 	Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.
//
// Example 2:
//
//	Input: 3
// 	Output: 2
// 	Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.
//
// Example 3:
//
//	Input: 4
// 	Output: 3
// 	Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.
//
//
// Note:
//
//	1. 0 ≤ N ≤ 30.
//
// https://leetcode.com/problems/fibonacci-number

// Using recursion.
func fibRecurse(n int) int {
	if n <= 1 {
		return n // base case: F(0)=0, F(1)=1
	}

	// Recurse for all other values of n.
	return fibRecurse(n-2) + fibRecurse(n-1)
}

// Use a lookup table to cache all numbers before 'n'. Thus, we don't
// have to compute the value for same numbers multiple times.
func fibLookup(n int) int {
	if n <= 1 {
		return n
	}

	// Cache of Fibonacci numbers for lookup.
	cache := make([]int, n+1)

	// First two Fibonacci numbers are 0 & 1.
	cache[0], cache[1] = 0, 1

	// Compute and save Fibonacci numbers upto n.
	// Use the cache for computatio.
	for i := 2; i <= n; i++ {
		cache[i] = cache[i-2] + cache[i-1]
	}

	return cache[n]
}

// Only last two Fibonacci numbers, n-2 and n-1, are required to compute
// n'th Fibonacci number. So, we just have to save last two Fibonacci
// numbers and not all.
func fib(n int) int {
	if n <= 1 {
		return n
	}

	f1, f2 := 0, 1 // base values: F(0)=0, F(1)=1
	for i := 2; i <= n; i++ {
		ith := f1 + f2 // i'th fib is the sum of previous two
		f1 = f2        // f[n-2]/f1 becomes f[n-1]/f2
		f2 = ith       // f2 is now i'th
	}

	// When the loop ends, f2 contins i'th number where i=n. So, f2
	// is the n'th Fibonacci number.
	return f2
}

func main() {
}
