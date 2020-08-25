package main

// Jewels and stones
//
// You're given strings J representing the types of stones that are
// jewels, and S representing the stones you have.  Each character in S
// is a type of stone you have.  You want to know how many of the stones
// you have are also jewels.
//
// The letters in J are guaranteed distinct, and all characters in J and
// S are letters. Letters are case sensitive, so "a" is considered a
// different type of stone from "A".
//
// Example:
//	Input: J = "aA", S = "aAAbbbb"
//	Output: 3
//
// Note:
//	- S and J will consist of letters and have length at most 50.
//	- The characters in J are distinct.
//
// https://leetcode.com/problems/jewels-and-stones/

func numJewelsInStones(jewels string, stones string) int {
	var stoneJewels = make(map[rune]bool)
	for _, j := range jewels {
		stoneJewels[j] = true
	}

	count := 0
	for _, stone := range stones {
		isJewel, present := stoneJewels[stone]
		if present && isJewel {
			count++
		}
	}

	return count
}

func main() {
}
