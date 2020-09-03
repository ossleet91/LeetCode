package main

// The K weakest rows in a matrix.
//
// Given a m * n matrix mat of ones (representing soldiers) and zeros
// (representing civilians), return the indexes of the k weakest rows in
// the matrix ordered from the weakest to the strongest.
//
// A row i is weaker than row j, if the number of soldiers in row i is
// less than the number of soldiers in row j, or they have the same
// number of soldiers but i is less than j. Soldiers are always stand in
// the frontier of a row, that is, always ones may appear first and then
// zeros.
//
// Example:
//
//	Input: mat =
//		[[1,1,0,0,0],
//		 [1,1,1,1,0],
//		 [1,0,0,0,0],
//		 [1,1,0,0,0],
//		 [1,1,1,1,1]],
//		k = 3
//	Output: [2,0,3]
//	Explanation:
//		The number of soldiers for each row is:
//			row 0 -> 2
//			row 1 -> 4
//			row 2 -> 1
//			row 3 -> 2
//			row 4 -> 5
//	Rows ordered from the weakest to the strongest are [2,0,3,1,4]
//
// https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix

import (
	"container/heap"
)

// Since arr is already sorted as soldiers (ones) followed by civilians
// (zeros), we can just use binary search to find the number of
// soldiers.
//
// Index at which civilians start is the number of soliders.
//
// E.g., if row ins {1, 1, 1, 0, 0, 0, 0}, civilian start at index 3 and
// number of soldiers is 3 as well.
func countSoldiers(row []int) int {
	tmp, lo, hi := 0, 0, len(row)
	for lo < hi {
		mid := (lo + hi) / 2
		if row[mid] > tmp {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return lo
}

// Following priority queue implementation is from an example in Go's
// documentation [1]
//
// Adjusted Item's memebers to reflect row/soldier terminology and
// modified Less() to suit our problem's constraint.
//
// [1]: https://golang.org/pkg/container/heap/#example__priorityQueue

type Item struct {
	row       int // heap value: soldier's row
	nsoldiers int // heap priority: soldier count for given row
	index     int // heap index: used by container/heap
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// If two rows have same # of soldiers, the row with smaller
	// index is considered weak.
	if pq[i].nsoldiers == pq[j].nsoldiers {
		return pq[i].row < pq[j].row
	}

	return pq[i].nsoldiers < pq[j].nsoldiers
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func kWeakestRows(mat [][]int, k int) []int {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// Push each row's index (value) and it's soldier count
	// (priority) in to a min-heap priority queue.
	for i, row := range mat {
		item := &Item{
			row:       i,
			nsoldiers: countSoldiers(row),
			index:     i,
		}

		heap.Push(&pq, item)
	}

	// Pop first 'k' nodes from priority queue. Since ous is a
	// min-heap, the first 'k' nodes is guarteed to have minimum
	// priority (i.e., weakest soliders in our terminology).
	weakest := make([]int, k)
	for i := 0; i < k && pq.Len() > 0; i++ {
		weakest[i] = heap.Pop(&pq).(*Item).row
	}

	return weakest
}

func main() {
}
