package main

// 448. Find All Numbers Disappeared in an Ar
//
// Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array),
// some elements appear twice and others appear once.
//
// Find all the elements of [1, n] inclusive that do not appear in this
// array.
//
// Could you do it without extra space and in O(n) runtime? You may
// assume the returned list does not count as extra space.
//
// Example:
//
//	Input: [4,3,2,7,8,2,3,1]
// 	Output: [5,6]
//
// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array

// Find missing numbers in a given array.
//
// Record the frequency of each number. Then, add every number with 0
// frequency to output.
//
// Time: O(N). Each element is visited at most once.
//
// Space: O(N). Additional space for freqCount.
func findDisappearedNumbersSpace(nums []int) []int {
	freqCount := make([]int, len(nums)+1) // +1 to account for 0
	for _, n := range nums {
		freqCount[n]++
	}

	// Start at 1 as problem description says min. element shall be
	// 1.
	var missing []int
	for i := 1; i < len(freqCount); i++ {
		if freqCount[i] == 0 {
			missing = append(missing, i)
		}
	}

	return missing
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Find missing numbers in a given array.
//
// Since the numbers in array are in range [1..len(arr)], we can use the
// elements to represent 'seen' attribute; i.e., scan the array and for
// every number that we see, store the negative value of the number in
// that index. At the end, any index positiion with non-negative values
// are numbers that we did not see, and thus, it forms our result.
//
// Time: O(N). Each element is visited at most once.
//
// Space: O(1). No additional space (other than result which problem
// says doesn't count against space).
//
// nums=[4 3 2 7 8 2 3 1]
// setting nums with negative values
//         i=0: v=3 nums=[4 3 2 -7 8 2 3 1]
//         i=1: v=2 nums=[4 3 -2 -7 8 2 3 1]
//         i=2: v=1 nums=[4 -3 -2 -7 8 2 3 1]
//         i=3: v=6 nums=[4 -3 -2 -7 8 2 -3 1]
//         i=4: v=7 nums=[4 -3 -2 -7 8 2 -3 -1]
//         i=5: v=1 nums=[4 -3 -2 -7 8 2 -3 -1]
//         i=6: v=2 nums=[4 -3 -2 -7 8 2 -3 -1]
//         i=7: v=0 nums=[-4 -3 -2 -7 8 2 -3 -1]
// creating result array
//         i=4 nums[4]=8 (+ve); append to missing=[5]
//         i=5 nums[5]=2 (+ve); append to missing=[5 6]
// returning missing=[5 6]
// ========
func findDisappearedNumbers(nums []int) []int {
	//fmt.Printf("nums=%v\n", nums)

	//fmt.Println("setting nums with negative values")
	for i := 0; i < len(nums); i++ {
		v := abs(nums[i]) - 1

		// Change only if we haven't changed the number to
		// negative already.
		if nums[v] > 0 {
			nums[v] = -nums[v]
		}
		//fmt.Printf("	i=%d: v=%d nums=%d\n", i, v, nums)
	}

	//fmt.Println("creating result array")
	var missing []int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			missing = append(missing, i+1)
			//fmt.Printf("	i=%d nums[%d]=%d (+ve); append to missing=%v\n", i, i, nums[i], missing)
		}
	}

	//fmt.Printf("returning missing=%v\n", missing)
	//fmt.Println("========")

	return missing
}

func main() {
}
