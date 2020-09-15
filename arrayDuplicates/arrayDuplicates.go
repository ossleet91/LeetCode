package main

// 442. Find All Duplicates in an Array
//
// Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array), some
// elements appear twice and others appear once.
//
// Find all the elements that appear twice in this array.
//
// Could you do it without extra space and in O(n) runtime?
//
// Example:
//
//	Input: [4,3,2,7,8,2,3,1]
//	Output: [2,3]
//
// https://leetcode.com/problems/find-all-duplicates-in-an-array

// Find numbers that appear twice in an array.
//
// Create a frequency of all numbers from [1..len(nums)]. Then, iterate
// the frequency and add any index with freq. = 2 to result array.
//
// Time: O(N). Each element in nums is visited at most once.
//
// Space: O(N). Additional array is allocated for freqency count.
//
// nums=[4 3 2 7 8 2 3 1]
// creatint freqCount
//         n=4 freqCount=[0 0 0 0 1 0 0 0 0]
//         n=3 freqCount=[0 0 0 1 1 0 0 0 0]
//         n=2 freqCount=[0 0 1 1 1 0 0 0 0]
//         n=7 freqCount=[0 0 1 1 1 0 0 1 0]
//         n=8 freqCount=[0 0 1 1 1 0 0 1 1]
//         n=2 freqCount=[0 0 2 1 1 0 0 1 1]
//         n=3 freqCount=[0 0 2 2 1 0 0 1 1]
//         n=1 freqCount=[0 1 2 2 1 0 0 1 1]
// creating result array
//         i=2 fc=2 (twice); append i to dups=[2]
//         i=3 fc=2 (twice); append i to dups=[2 3]
// returning dups=[2 3]
// ========.
func findDuplicatesSpace(nums []int) []int {
	//fmt.Printf("nums=%v\n", nums)

	//fmt.Println("creatint freqCount")
	freqCount := make([]int, len(nums)+1)
	for _, n := range nums {
		freqCount[n]++
		//fmt.Printf("	n=%d freqCount=%v\n", n, freqCount)
	}

	//fmt.Println("creating result array")
	var dups []int
	for i, fc := range freqCount {
		if fc == 2 {
			dups = append(dups, i)
			//fmt.Printf("	i=%d fc=%d (twice); append i to dups=%v\n", i, fc, dups)
		}
	}

	//fmt.Printf("returning dups=%v\n", dups)
	//fmt.Println("========.")

	return dups
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Find numbers that appear twice in an array.
//
// Iterate through the array and convert every element to index (i.e.,
// element-1) and flip the sign of element at that index. Before
// flipping, if the element in that position is negative, convert the
// index to element (i.e., index+1) and add it to result array.
//
// This works because problem guarantees that elements are in the range
// [1..len(nums)]. So, if there's a duplicate element, it's index
// postion would have been negated the first time around. So, when the
// same element appears again, it will be negative and can be added to
// the list.
//
// Time: O(N). Each element in nums is visited at most once.
//
// Space: O(1). No additional space, except for result array, is used.
// Problem says result array don't count for space complexity.
//
// nums=[4 3 2 7 8 2 3 1]
// marking visited as negative
//         i=0 n=4 index=3:
//                 negated nums[index=3]=-7; nums=[4 3 2 -7 8 2 3 1]
//         i=1 n=3 index=2:
//                 negated nums[index=2]=-2; nums=[4 3 -2 -7 8 2 3 1]
//         i=2 n=-2 index=1:
//                 negated nums[index=1]=-3; nums=[4 -3 -2 -7 8 2 3 1]
//         i=3 n=-7 index=6:
//                 negated nums[index=6]=-3; nums=[4 -3 -2 -7 8 2 -3 1]
//         i=4 n=8 index=7:
//                 negated nums[index=7]=-1; nums=[4 -3 -2 -7 8 2 -3 -1]
//         i=5 n=2 index=1:
//                 nums[index=1]=-3 (-ve); so append index+1=2 to dups=[2]
//                 negated nums[index=1]=3; nums=[4 3 -2 -7 8 2 -3 -1]
//         i=6 n=-3 index=2:
//                 nums[index=2]=-2 (-ve); so append index+1=3 to dups=[2 3]
//                 negated nums[index=2]=2; nums=[4 3 2 -7 8 2 -3 -1]
//         i=7 n=-1 index=0:
//                 negated nums[index=0]=-4; nums=[-4 3 2 -7 8 2 -3 -1]
// returning dups=[2 3]
// ========.
func findDuplicates(nums []int) []int {
	//fmt.Printf("nums=%v\n", nums)

	//fmt.Println("marking visited as negative")

	var dups []int
	for _, n := range nums {
		index := abs(n) - 1
		//fmt.Printf("	i=%d n=%d index=%d:\n", i, n, index)

		if nums[index] < 0 {
			dups = append(dups, index+1)
			//fmt.Printf("		nums[index=%d]=%d (-ve); so append index+1=%d to dups=%v\n", index, nums[index], index+1, dups)
		}

		nums[index] = -nums[index]
		//fmt.Printf("		negated nums[index=%d]=%d; nums=%v\n", index, nums[index], nums)
	}

	//fmt.Printf("returning dups=%v\n", dups)
	//fmt.Println("========.")

	return dups
}

func main() {
}
