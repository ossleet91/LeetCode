package main

// Number of students doing homework at a given time.
//
// Given two integer arrays startTime and endTime and given an integer queryTime.
//
// The ith student started doing their homework at the time startTime[i]
// and finished it at time endTime[i].
//
// Return the number of students doing their homework at time queryTime.
// More formally, return the number of students where queryTime lays in
// the interval [startTime[i], endTime[i]] inclusive.
//
// Example:
//
//	Input: startTime = [1,2,3], endTime = [3,2,7], queryTime = 4
//	Output: 1
//	Explanation: We have 3 students where:
//	The first student started doing homework at time 1 and finished
//	at time 3 and wasn't doing anything at time 4.  The second
//	student started doing homework at time 2 and finished at time 2
//	and also wasn't doing anything at time 4.  The third student
//	started doing homework at time 3 and finished at time 7 and was
//	the only student doing homework at time 4.
//
// https://leetcode.com/problems/number-of-students-doing-homework-at-a-given-time/

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	busy := 0

	for i := 0; i < len(startTime); i++ {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			busy++
		}
	}

	return busy
}

func main() {
}
