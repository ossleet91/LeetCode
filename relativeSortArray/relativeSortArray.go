package main

// Relative sort array
//
// Given two arrays arr1 and arr2, the elements of arr2 are distinct,
// and all elements in arr2 are also in arr1.
//
// Sort the elements of arr1 such that the relative ordering of items in
// arr1 are the same as in arr2.  Elements that don't appear in arr2
// should be
// placed at the end of arr1 in ascending order.
//
//Example:
//
//	Input: arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
//	Output: [2,2,2,1,4,3,3,9,6,7,19]
//
//Constraints:
//
//	1. arr1.length, arr2.length <= 1000
//	2. 0 <= arr1[i], arr2[i] <= 1000
//	3. Each arr2[i] is distinct.
//	4. Each arr2[i] is in arr1.
//
// https://leetcode.com/problems/relative-sort-array

import (
	"sort"
)

// CountMap is a map of value to count.
//
// This is useful when we have to know how many times a number is
// present in an array.
//
// E.g., []int{1, 1, 1, 3, 1, 2}'s CountMap will be {1:4, 3:1, 2:1}.
type CountMap map[int]int

// Create and return a new CountMap for given arr.
func newCountMapValue(arr []int) CountMap {
	cm := CountMap{}

	for _, n := range arr {
		cm.Add(n)
	}

	return cm
}

// Add given number to CountMap.
func (cm CountMap) Add(n int) {
	cm[n]++
}

// Delete given number from CountMap. If the count is 1, delete the
// number from map. Otherwise, decrement its count.
func (cm CountMap) Del(n int) {
	val, ok := cm[n]
	if ok {
		if val > 1 {
			cm[n]--
		} else {
			delete(cm, n)
		}
	}
}

// Delete given number from CountMap.
func (cm CountMap) DelFull(n int) {
	delete(cm, n)
}

// Check if a given number is present in CountMap.
func (cm CountMap) Present(n int) bool {
	_, ok := cm[n]
	return ok
}

// Using CountMap.
func relativeSortArray(arr1, arr2 []int) []int {
	cm1 := newCountMapValue(arr1)

	sorted := make([]int, 0)
	for _, n := range arr2 {
		for i := 0; i < cm1[n]; i++ {
			sorted = append(sorted, n)
		}
		cm1.DelFull(n)
	}

	extra := make([]int, 0)
	for n, count := range cm1 {
		for i := 0; i < count; i++ {
			extra = append(extra, n)
		}
		cm1.DelFull(n)
	}

	sort.Ints(extra)
	return append(sorted, extra...)
}

// Using CountMap, less space.
//
// Time: O(n log_n). O(n) for preparing CountMap, O(n-k log_n-k) for
// sorting elements in arr1 that are not in arr2.
//
// Space: O(n). O(n) for CountMap. We reuse arr1 for result.

func relativeSortArrayLessSpace(arr1, arr2 []int) []int {
	cm1 := newCountMapValue(arr1)

	i := 0

	// Store every element in arr2 based on it's arr1 count from
	// CountMap.
	for _, n := range arr2 {
		count := cm1[n]
		for count > 0 {
			arr1[i] = n
			i++
			count--
		}
		cm1.DelFull(n)
	}

	// Store remaining elements (not in arr2, but in arr1) from
	// CountMap. Then, sort just this portion.
	extraIndex := i
	for n, count := range cm1 {
		for count > 0 {
			arr1[i] = n
			i++
			count--
		}
		cm1.DelFull(n)
	}
	sort.Ints(arr1[extraIndex:])

	return arr1
}

// Using couting sort.
func relativeSortArrayCS(arr1, arr2 []int) []int {
	counter := make([]int, 1001) // see constraint #2
	for _, n := range arr1 {
		counter[n]++
	}

	i := 0
	for _, n := range arr2 {
		for counter[n] > 0 {
			arr1[i] = n
			counter[n]--
			i++
		}
	}

	for n := range counter {
		for counter[n] > 0 {
			arr1[i] = n
			counter[n]--
			i++
		}
	}

	return arr1
}

func main() {
}
