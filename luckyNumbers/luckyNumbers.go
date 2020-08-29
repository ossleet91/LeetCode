package main

// Lucky numbers in a marrix.
//
// Given a m * n matrix of distinct numbers, return all lucky numbers in
// the matrix in any order.
//
// A lucky number is an element of the matrix such that it is the
// minimum element in its row and maximum in its column.
//
//
//
// Example:
//
//	Input: matrix = [[3,7,8],[9,11,13],[15,16,17]]
//	Output: [15]
//	Explanation: 15 is the only lucky number since it is the minimum
//	in its row and the maximum i
//
// Constraints:
//
//	1. m == mat.length
//	2. n == mat[i].length
//	3. 1 <= n, m <= 50
//	4. 1 <= matrix[i][j] <= 10^5.
//	5. All elements in the matrix are distinct.
//
// https://leetcode.com/problems/lucky-numbers-in-a-matrix/

func minInArr(arr []int) int {
	min := 100_000 // constraint #4 says max value is 10^5
	for _, n := range arr {
		if n < min {
			min = n
		}
	}

	return min
}

func maxInArr(arr []int) int {
	max := 1 // constraint #4 says min value is 1
	for _, n := range arr {
		if n > max {
			max = n
		}
	}

	return max
}

func getColumn(matrix [][]int, colIndex int) []int {
	column := make([]int, len(matrix))
	for i := range matrix {
		column[i] = matrix[i][colIndex]
	}

	return column
}

func luckyNumbers(matrix [][]int) []int {
	nrows, ncols := len(matrix), len(matrix[0])

	minrow := make([]int, nrows)
	maxcol := make([]int, ncols)

	for i := 0; i < nrows; i++ {
		minrow[i] = minInArr(matrix[i])
	}

	for i := 0; i < ncols; i++ {
		col := getColumn(matrix, i)
		maxcol[i] = maxInArr(col)
	}

	lucky := make([]int, 0)
	for r := 0; r < nrows; r++ {
		for c := 0; c < ncols; c++ {
			n := matrix[r][c]
			if n == minrow[r] && n == maxcol[c] {
				lucky = append(lucky, n)
			}
		}
	}

	return lucky
}

func main() {
}
