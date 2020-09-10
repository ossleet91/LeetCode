package main

// Remove a given element in place.
//
// Given an array nums and a value val, remove all instances of that
// value in-place and return the new length.
//
// Do not allocate extra space for another array, you must do this by
// modifying the input array in-place with O(1) extra memory.
//
// The order of elements can be changed. It doesn't matter what you
// leave beyond the new length.
//
// Example:
//
// Given nums = [3,2,2,3], val = 3,
//
// Your function should return length = 2, with the first two elements
// of nums being 2.
//
// It doesn't matter what you leave beyond the returned length.
//
// https://leetcode.com/problems/remove-element/

// Count the number of times 'val' appears in 'nums'.
func countOf(nums []int, val int) int {
	count := 0

	for _, n := range nums {
		if n == val {
			count++
		}
	}

	return count
}

// Remove all occurances of a given element from the array and re-order
// the remaning elements in-place. Return the length of new array.
//
// Use two-pointers approach. Forward pointer scans the whole array for
// element to be removed. Back pointers is from the back and is used for
// storing elements that needs to be discarded.
//
// Time: O(N). All elements are visited at most once.
//
// Space: O(1). Just two pointers and no additional array.
//
// nums=[0 1 2 2 3 0 4 2] val=2
//   fwd/n[f]=0/0 bck/n[b]=7/2 nums=[0 1 2 2 3 0 4 2]
//   fwd/n[f]=1/1 bck/n[b]=7/2 nums=[0 1 2 2 3 0 4 2] // n[f]==val; swap with n[b] after finding correct b in next step
//   fwd/n[f]=2/4 bck/n[b]=6/2 nums=[0 1 4 2 3 0 2 2] // b decremented to 6 as n[b]==val and swapped with f
//   fwd/n[f]=3/0 bck/n[b]=5/2 nums=[0 1 4 0 3 2 2 2] // same as above
//   fwd/n[f]=4/3 bck/n[b]=5/2 nums=[0 1 4 0 3 2 2 2]
//   fwd/n[f]=5/2 bck/n[b]=5/2 nums=[0 1 4 0 3 2 2 2]
//   fwd/n[f]=6/2 bck/n[b]=5/2 nums=[0 1 4 0 3 2 2 2]
//   fwd/n[f]=7/2 bck/n[b]=5/2 nums=[0 1 4 0 3 2 2 2] // done; b points to final array with all val removed.
func removeElement(nums []int, val int) int {
	valCount := countOf(nums, val)
	if valCount == 0 {
		return len(nums)
	} else if valCount == len(nums) {
		return 0
	}

	bck := len(nums) - 1
	for fwd := 0; fwd < len(nums); fwd++ {
		// Find the first position from back whose value isn't
		// same as what needs to be removed.
		//
		// Then, swap it with forward's value which needs to be
		// removed.
		//
		// This guarantees all elements from back till the end
		// can be discarded.
		if nums[fwd] == val {
			for nums[fwd] == nums[bck] && bck > fwd {
				bck--
			}
			nums[fwd], nums[bck] = nums[bck], nums[fwd]
		}
	}

	return bck
}

// Remove all occurances of a given element from the array and re-order
// the remaning elements in-place. Return the length of new array.
//
// Use two-pointers approach. Fast and slow starts at 0. Fast increments
// at every iteration. If 'nums[fast] != val', copy slow to fast and
// increment slow. Do not increment slow when fast points to 'val'. This
// guarantees slow values (the ones that needs to be removed) will be
// copied over with fast values. At the end, slow points to end of all
// 'val' removed array.
//
// Time: O(N). All elements are visited at most once.
//
// Space: O(1). Just two pointers and no additional array.
//
// nums=[]int{0, 1, 2, 2, 3, 0, 4, 2} val=2
//   slow=1 fast=0 nums=[]int{0, 1, 2, 2, 3, 0, 4, 2}
//   slow=2 fast=1 nums=[]int{0, 1, 2, 2, 3, 0, 4, 2}
//   slow=2 fast=2 nums=[]int{0, 1, 2, 2, 3, 0, 4, 2} // s remains at 2 since nums[2]==val
//   slow=2 fast=3 nums=[]int{0, 1, 2, 2, 3, 0, 4, 2} // s remains at 2 since nums[2]==val
//   slow=3 fast=4 nums=[]int{0, 1, 3, 2, 3, 0, 4, 2} // f moved to non-val and we copy f to s
//   slow=4 fast=5 nums=[]int{0, 1, 3, 0, 3, 0, 4, 2}
//   slow=5 fast=6 nums=[]int{0, 1, 3, 0, 4, 0, 4, 2}
//   slow=5 fast=7 nums=[]int{0, 1, 3, 0, 4, 0, 4, 2} // done; s points to len of final array.
func removeElement1(nums []int, val int) int {
	fast, slow := 0, 0
	for ; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}

// Remove all occurances of a given element from the array and re-order
// the remaning elements in-place. Return the length of new array.
//
// Use two-pointers approach. i and n starts at 0 and len(nums). When
// n[i] == val, swap it with n-1 and decrement back pointer. Otherwise,
// increment n. This works as we don't have to maintain order in the
// resulting array.
//
// Time: O(N). All elements are visited at most once.
//
// Space: O(1). Just two pointers and no additional array.
// nums=[0 1 2 2 3 0 4 2] val=2
//   i/n[i]=1/1 n/n[n-1]=8/2 nums=[0 1 2 2 3 0 4 2]
//   i/n[i]=2/2 n/n[n-1]=8/2 nums=[0 1 2 2 3 0 4 2] // swap i & n in next step; decrement n
//   i/n[i]=2/2 n/n[n-1]=7/4 nums=[0 1 2 2 3 0 4 2] // swap i & n in next step; decrement n
//   i/n[i]=2/4 n/n[n-1]=6/0 nums=[0 1 4 2 3 0 4 2]
//   i/n[i]=3/2 n/n[n-1]=6/0 nums=[0 1 4 2 3 0 4 2] // swap i & n in next step; decrement n
//   i/n[i]=3/0 n/n[n-1]=5/3 nums=[0 1 4 0 3 0 4 2]
//   i/n[i]=4/3 n/n[n-1]=5/3 nums=[0 1 4 0 3 0 4 2]
//   i/n[i]=5/0 n/n[n-1]=5/3 nums=[0 1 4 0 3 0 4 2] // done; n points to length of final array.
func removeElement2(nums []int, val int) int {
	i, n := 0, len(nums)
	for i < n {
		if nums[i] == val {
			nums[i] = nums[n-1]
			n--
		} else {
			i++
		}
	}

	return n
}

func main() {
}
