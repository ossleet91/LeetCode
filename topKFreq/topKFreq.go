package main

// Top K frequent elements
//
// Given a non-empty array of integers, return the k most frequent
// elements.
//
// Example:
//	Input: nums = [1,1,1,2,2,3], k = 2
//	Output: [1,2]
//
// Note:
//	1. You may assume k is always valid, 1 ≤ k ≤ number of unique
//	   elements.
//
//	2. Your algorithm's time complexity must be better than
//	   O(n logn), where n is the array's size.
//
//	3. It's guaranteed that the answer is unique, in other words the
//	   set of the top k frequent elements is unique.
//
//	4. You can return the answer in any order.
//
// https://leetcode.com/problems/top-k-frequent-elements

import (
	"container/heap"
)

func topKFreq(nums []int, k int) []int {
	// Creata a map of each number and its frequency.
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}

	// Swap the freq map's key and values.
	// i.e., map[number]=count becomes map[count]=[]int{number...}
	freqVal := make(map[int][]int)
	for n, count := range freq {
		freqVal[count] = append(freqVal[count], n)
	}

	// If there are len(nums) values in nums, then, there can only
	// be len(nums) frequencies at most (consider same number
	// repeated all over).
	//
	// So, go over all possible frequencies in decreasing order. If
	// ietrating frequecy is present in our map, append map[freq] to
	// result.
	var res []int
	for f := len(nums); f > 0; f-- {
		if n, ok := freqVal[f]; ok {
			res = append(res, n...)
		}

		// Stop once we accumulate at least 'k' elements.
		// > k because res could grow beyond k due to slice
		// concat.
		if len(res) >= k {
			return res[:k]
		}
	}

	return res
}

type Item struct {
	n     int // value to track in heap
	count int // n's frequency or count
	index int // required by container/heap
}

type PriorityQueue []*Item

func (pq *PriorityQueue) Len() int {
	return len(*pq)
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.index = pq.Len()
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := pq.Len()
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[:n-1]
	return item
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].count > pq[j].count
}

func topKFreqPQ(nums []int, k int) []int {
	// Creata a map of each number and its frequency.
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	i := 0
	for n, count := range freq {
		item := Item{
			n:     n,
			count: count,
			index: i,
		}
		heap.Push(&pq, &item)
		i++
	}

	var res []int
	for i := 0; i < k && pq.Len() > 0; i++ {
		item := heap.Pop(&pq).(*Item)
		res = append(res, item.n)
	}
	return res
}

func main() {
}
