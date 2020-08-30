package main

// Average salary excluding minimum and maximum salary.
//
// Given an array of unique integers salary where salary[i] is the salary of the employee i.
//
// Return the average salary of employees excluding the minimum and maximum salary.
//
// Example:
//
//	Input: salary = [4000,3000,1000,2000]
//	Output: 2500.00000
//	Explanation: Minimum salary and maximum salary are 1000 and 4000
//		     respectively. Average salary excluding minimum and maximum
//		     salary is (2000+3000)/2= 2500
//
// Constraints:
//
//	1. 3 <= salary.length <= 100
//	2. 10^3 <= salary[i] <= 10^6
//	3. salary[i] is unique.
//	4. Answers within 10^-5 of the actual value will be accepted as correct.
//
// https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary

func averageSalary(salary []int) float64 {
	// Constraint #2 limits min and max values to 10^3 and 10^6.
	min, max := uint64(1_000_000), uint64(1_000)

	var sum uint64
	for i := range salary {
		s := uint64(salary[i])

		if s < min {
			min = s
		}

		if s > max {
			max = s
		}

		sum += s
	}

	return float64(sum-min-max) / float64(len(salary)-2)
}

func main() {
}
