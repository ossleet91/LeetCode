package main

// Remove duplicates from sorted array.
//
// Given a sorted array nums, remove the duplicates in-place such that each element appear only once and return the new length.
//
// Do not allocate extra space for another array, you must do this by
// modifying the input array in-place with O(1) extra memory.
//
// Example 1
//
//	Given nums = [1,1,2],
//
//	Your function should return length = 2, with the first two
//	elements of nums being 1 and 2 respectively.
//
//	It doesn't matter what you leave beyond the returned length.
//
// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

// Returns count of unique elements in a given sorted slice 'nums'.
// 'max' represents the largest value in 'nums'.
func unique(nums []int, max int) int {
	var count, prev int
	for _, n := range nums {
		if n == max {
			count++
			break
		}

		if n == prev {
			continue
		}

		prev = n
		count++
	}

	return count
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

// Rotate 'array' by 'count' positions.
// rotate([]int{1, 1, 3, 5, 6}, 2) = {3, 5, 6, 1 ,1}
func rotateLeft(array []int, count int) {
	count %= len(array)
	gcd := gcd(count, len(array))

	for i := 0; i < gcd; i++ {
		tmp := array[i]
		j := i

		for {
			k := j + count
			if k >= len(array) {
				k -= len(array)
			}

			if k == i {
				break
			}

			array[j] = array[k]
			j = k

		}
		array[j] = tmp
	}
}

func rotateLeftOne(array []int, count int) {
	if count == 0 || count == len(array) {
		return
	}

	var i, j int
	for i = 0; i < count; i++ {
		tmp := array[0]
		for j = 0; j < len(array)-1; j++ {
			array[j] = array[j+1]
		}
		array[j] = tmp
	}
}

func rotateLeftTmp(array []int, count int) {
	if count == 0 || count == len(array) {
		return
	}

	tmp := make([]int, count)
	for i := range tmp {
		tmp[i] = array[i]
	}

	forward := len(array) - count
	backward := count

	for i := 0; i < forward; i++ {
		array[i] = array[i+count]
	}

	for i := 0; i < backward; i++ {
		array[i+forward] = tmp[i]
	}
}

func removeDuplicates1(nums []int) int {
	var i, j int

	for j = 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}

func removeDuplicates(nums []int) int {
	var prev, nMoved, dupStart, dupEnd int
	var procDups bool

	max := nums[len(nums)-1]
	if nums[0] == max {
		// Return soon if the first element itself is the
		// largest.  Since 'nums' is sorted, this ensures all
		// the elements are same as first element.
		return 1
	}

	for i := 0; i < len(nums); i++ {
		curr := nums[i]

		if i == 0 {
			prev = curr
			continue
		}

		if curr == prev {
			if !procDups {
				dupStart = i
				procDups = true
			}
		} else if curr > prev {
			if procDups {
				dupEnd = i - 1
				procDups = false

				// Now we have identified a set of dups.
				// Just rotate the array so that dups go
				// to the end.
				count := dupEnd - dupStart + 1

				// Select elements only from dupStart to
				// end of non-visited elements.
				//
				// i.e., don't bother with duplicate
				// elements that we moved to the end.
				arr := nums[dupStart : len(nums)-nMoved]
				rotateLeft(arr, count)

				// Keep track of moved elments so that
				// we don't have to include them in
				// rotation next time around.
				nMoved += count

				// Reset array index so that we start
				// from our 'curr' element again.
				i -= count
			}
		}

		if curr == max {
			// Current element is the largest element and
			// our running index has been adjusted to
			// account for removed duplicates. So, [0..i] is
			// our current result. Return now.
			return i + 1
		}

		prev = curr
	}

	return len(nums) - nMoved
}

func doRotate() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	count := 3

	rotateLeftTmp(arr, count)
}

func main() {
}
