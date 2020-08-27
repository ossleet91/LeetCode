package main

// Minimum time to visit all the points
//
// On a plane there are n points with integer coordinates points[i] =
// [xi, yi].  Your task is to find the minimum time in seconds to visit
// all points.
//
// You can move according to the next rules:
//
//	- In one second always you can either move vertically,
//	  horizontally by one unit or diagonally (it means to move one
//	  unit vertically and one unit horizontally in one second).
//	- You have to visit the points in the same order as they appear
//	  in the array.
//
// Example:
//	Input: points = [[1,1],[3,4],[-1,0]]
//	Output: 7
//	Explanation:
//		One optimal path is [1,1] -> [2,2] -> [3,3] -> [3,4]
//				 -> [2,3] -> [1,2] -> [0,1] -> [-1,0]
//		Time from [1,1] to [3,4] = 3 seconds
//		Time from [3,4] to [-1,0] = 4 seconds
//		Total time = 7 seconds
//
// https://leetcode.com/problems/minimum-time-visiting-all-points/

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func minTimeToVisitAllPoints(points [][]int) int {
	steps := 0

	for i := 0; i < len(points)-1; i++ {
		x, y := points[i][0], points[i][1]
		dx, dy := points[i+1][0], points[i+1][1]
		steps += max(abs(dx-x), abs(dy-y))
	}

	return steps
}

func step(val, dest int) int {
	if val < dest {
		return val + 1
	} else if val > dest {
		return val - 1
	}

	return val
}

func minTimeToVisitAllPointsSlow(points [][]int) int {
	steps := 0

	x, y := points[0][0], points[0][1]
	for i := 1; i < len(points); i++ {
		dx, dy := points[i][0], points[i][1]

		for x != dx || y != dy {
			if x != dx {
				x = step(x, dx)
			}

			if y != dy {
				y = step(y, dy)
			}

			steps++
		}
	}

	return steps
}

func main() {
}
