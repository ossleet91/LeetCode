package main

// Cells with odd values in a matrix
//
// Given n and m which are the dimensions of a matrix initialized by
// zeros and given an array indices where indices[i] = [ri, ci]. For
// each pair of [ri, ci] you have to increment all cells in row ri and
// column ci by 1.
//
// Return the number of cells with odd values in the matrix after
// applying the increment to all indices.
//
// https://leetcode.com/problems/cells-with-odd-values-in-a-matrix/

func countOddCells(n, m int, indices [][]int) int {
	mat := make([][]int, n)
	for i := range mat {
		mat[i] = make([]int, m)
	}

	for i := 0; i < len(indices); i++ {
		row, col := indices[i][0], indices[i][1]

		for r := 0; r < m; r++ {
			mat[row][r]++
		}

		for c := 0; c < n; c++ {
			mat[c][col]++
		}
	}

	odd := 0
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if mat[r][c]%2 != 0 {
				odd++
			}
		}
	}

	return odd
}

func main() {
}
