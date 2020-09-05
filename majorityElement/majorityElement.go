package main

// 169. Majority Element
//
// Given an array of size n, find the majority element. The majority
// element is the element that appears more than ⌊ n/2 ⌋ times.
//
// You may assume that the array is non-empty and the majority element
// always exist in the array.
//
// Example 1:
//
//	Input: [3,2,3]
// 	Output: 3
//
// Example 2:
//
//	Input: [2,2,1,1,1,2,2]
// 	Output: 2
//
// https://leetcode.com/problems/majority-element

type CountMap map[int]int

func newCountMap(arr []int) CountMap {
	cm := CountMap{}

	for _, n := range arr {
		cm.Add(n)
	}

	return cm
}

func (cm CountMap) Add(n int) {
	cm[n]++
}

func (cm CountMap) Del(n int) {
	if val, ok := cm[n]; ok && val > 1 {
		cm[n]--
	} else {
		delete(cm, n)
	}
}

func (cm CountMap) Present(n int) bool {
	_, ok := cm[n]
	return ok
}

func (cm CountMap) MaxCount() int {
	maxN, maxCount := 0, 0
	for n, count := range cm {
		if count > maxCount {
			maxCount = count
			maxN = n
		}
	}

	return maxN
}

func majorityElement(nums []int) int {
	cm := newCountMap(nums)
	return cm.MaxCount()
}

// Find majority element using Boyer-Moore majority voting algorithm.
// https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore_majority_vote_algorithm
func majorityElementBoyerMoore(nums []int) int {
	count, candidate := 0, 0
	for _, n := range nums {
		switch {
		case count == 0:
			candidate = n
			count = 1
		case n == candidate:
			count++
		default:
			count--
		}
	}

	// Boyer-Moore voting alogorithm finds majority if one exists or
	// reports no-majority. This is accomplished by finding a
	// majoirty first and then, in a second pass over the array,
	// compute the frequency of the just-found-candidate. If the
	// frequency > len(arr)/2, return candidate as majjority.
	// Otherwise, there's no majority element in the given array.
	//
	//	majCount := 0
	// 	for _, n := range nums {
	// 		if n == candidate {
	// 			majCount++
	// 		}
	// 	}
	//
	// 	if majCount > len(arr)/2 {
	// 		// Majority element present.
	// 		return true, candidate
	// 	}
	// 	// No element has majority.
	// 	return false, nil
	//
	// But our problem guarantees that at leaset one element's
	// frequency is more than len(arr)/2. So, we can skip the
	// second-pass and return our candidate as the majoirty-element.

	return candidate
}

func main() {
}
