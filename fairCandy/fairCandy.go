package main

// 888. Fair Candy Swap
//
// Alice and Bob have candy bars of different sizes: A[i] is the size of
// the i-th bar of candy that Alice has, and B[j] is the size of the
// j-th bar of candy that Bob has.
//
// Since they are friends, they would like to exchange one candy bar
// each so that after the exchange, they both have the same total amount
// of candy.  (The total amount of candy a person has is the sum of the
// sizes of candy bars they have.)
//
// Return an integer array ans where ans[0] is the size of the candy bar
// that Alice must exchange, and ans[1] is the size of the candy bar
// that Bob must exchange.
//
// If there are multiple answers, you may return any one of them.  It is
// guaranteed an answer exists.
//
//
// Example 1:
//
//	Input: A = [1,1], B = [2,2]
// 	Output: [1,2]
//
// Example 2:
//
//	Input: A = [1,2], B = [2,3]
// 	Output: [1,2]
//
// Example 3:
//
//	Input: A = [2], B = [1,3]
// 	Output: [2,3]
//
// Example 4:
//
//	Input: A = [1,2,5], B = [2,4]
// 	Output: [5,4]
//
// Note:
//
//	1. 1 <= A.length <= 10000
//	2. 1 <= B.length <= 10000
//	3. 1 <= A[i] <= 100000
//	4. 1 <= B[i] <= 100000
//	5. It is guaranteed that Alice and Bob have different total
//	   amounts of candy.
//	6. It is guaranteed there exists an answer.
//
// https://leetcode.com/problems/fair-candy-swap

type HashSet map[int]bool

func newHashSet(arr []int) HashSet {
	hs := HashSet{}
	for _, n := range arr {
		hs[n] = true
	}

	return hs
}

func (hs HashSet) Present(n int) bool {
	_, ok := hs[n]
	return ok
}

func sum(nums []int) int {
	var sum uint64
	for _, n := range nums {
		sum += uint64(n)
	}

	return int(sum)
}

// To ensure both, Alice & Bob, has equal sized candies, we have ensure:
//
//	aShare - aXchange + bXchange = bShare - bXchange + aXchange
//	sShare - bShare = 2 * aXchange-bXchange
//	aShare-bShare + 2*bXchange = 2*aXchange
//	(aShare-bShare / 2) + bXchange = aXchange
//
// This is fairly easy when a HashSet is used.
func fairCandySwap(alice, bob []int) []int {
	aliceSum, bobSum := sum(alice), sum(bob)
	aliceSet := newHashSet(alice)
	diff := (aliceSum - bobSum) / 2

	for _, b := range bob {
		if aliceSet.Present(b + diff) {
			return []int{diff + b, b}
		}
	}

	return []int{0, 0}
}

func main() {
}
