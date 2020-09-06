package main

// 1395. Count Number of Teams
//
// There are n soldiers standing in a line. Each soldier is assigned a
// unique rating value.
//
// You have to form a team of 3 soldiers amongst them under the
// following rules:
//
//     - Choose 3 soldiers with index (i, j, k) with rating:
//	 (rating[i], rating[j], rating[k]).

//     - A team is valid if:  (rating[i] < rating[j] < rating[k]) or
//       (rating[i] > rating[j] > rating[k]) where (0 <= i < j < k < n).
//
// Return the number of teams you can form given the conditions.
// (soldiers can be part of multiple teams).
//
// Example 1:
//
//	Input: rating = [2,5,3,4,1]
// 	Output: 3
// 	Explanation: We can form three teams given the conditions.
// 	             (2,3,4), (5,4,1), (5,3,1).
//
// Example 2:
//
//	Input: rating = [2,1,3]
// 	Output: 0
// 	Explanation: We can't form any team given the conditions.
//
// Example 3:
//
//	Input: rating = [1,2,3,4]
// 	Output: 4
//
//
//
// Constraints:
//
//	1. n == rating.length
//	2. 1 <= n <= 200
//	3. 1 <= rating[i] <= 10^5
//
// https://leetcode.com/problems/count-number-of-teams

// Naive approach.
//
// Select triplets {i,i+1,i+2} for i=0..len-2 and for each each triplet,
// increment a running counter by 1 if the problem's required condition
// holds.
//
// Time: O(N^3). Three nested loops, each going over array.
// Space: O(1). Just a single running counter.
func numTeamsNaive(ratings []int) int {
	res := 0

	if len(ratings) < 3 {
		return res
	}

	// Limit to len-2 to prevent out-of-bound array access as final
	// loop is essentially i+2 (or j+1).
	for i := 0; i < len(ratings)-2; i++ {
		// Limit to len-1 to prevent out-of-bound array access
		// as final loop is j+1.
		for j := i + 1; j < len(ratings)-1; j++ {
			for k := j + 1; k < len(ratings); k++ {
				// Don't have to worry about <= or >= as
				// problem guarantees every rating is
				// unique.
				ri, rj, rk := ratings[i], ratings[j], ratings[k]
				if ri < rj && rj < rk || ri > rj && rj > rk {
					res++
				}
			}
		}
	}

	return res
}

// Mid-rating approach.
//
// Consider each rating as middle-rating in the triplet. Then, based on
// the problem's constraint, we can consider the triplet as valid only
// if the left and right ratings of mid are smaller and bigger than mid
// (or vice-versa).
//
// For each rating, mid, accumulate the count of smaller/bigger ratings
// to mid's left and right as leftSmall, leftBig, rightSmall, rightBig.
//
// Then, no. of valid triplets that can be formed with 'mid' in the
// middle is:
//	leftSmall*rightBig (ri < rj < rk) +
//	rightSmall*leftBig (ri > rj > rk)
//
// Compute this for every rating and accumulate the valid triplets as
// result.
//
// Time: O(N^2). Outer-loop with each element as mid and inner-loop to
//	 find smaller/bigger ratings.
// Space: O(1). Just five counters (one each for left/right small/big
//        and result).
func numTeams(ratings []int) int {
	res := 0

	if len(ratings) < 3 {
		return res
	}

	// We can just consider ratings[1..n-1] as first/last elements,
	// by definition, cannot be in the middle of a triplet.
	for i := 1; i < len(ratings)-1; i++ {
		mid := ratings[i]
		smallerLeft, biggerLeft := 0, 0
		smallerRight, biggerRight := 0, 0

		// Scan left-side for smaller/bigger ratings than mid.
		for j := i - 1; j >= 0; j-- {
			if ratings[j] < mid {
				smallerLeft++
			}
			if ratings[j] > mid {
				biggerLeft++
			}
		}

		// Scan right-side for smaller/bigger ratings than mid.
		for j := i + 1; j < len(ratings); j++ {
			right := ratings[j]
			if right < mid {
				smallerRight++
			}
			if right > mid {
				biggerRight++
			}
		}

		res += smallerLeft*biggerRight + smallerRight*biggerLeft
	}

	return res
}

func main() {
}
