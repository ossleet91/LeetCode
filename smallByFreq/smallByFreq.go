package main

// 1170. Compare Strings by Frequency of the Smallest Character
//
//Let's define a function f(s) over a non-empty string s, which
//calculates the frequency of the smallest character in s. For example,
//if s = "dcce" then f(s) = 2 because the smallest character is "c" and
//its frequency is 2.
//
// Now, given string arrays queries and words, return an integer array
// answer, where each answer[i] is the number of words such that
// f(queries[i]) < f(W), where W is a word in words.
//
//
//
// Example 1:
//
//	Input: queries = ["cbd"], words = ["zaaaz"]
// 	Output: [1]
// 	Explanation: On the first query we have f("cbd") = 1, f("zaaaz")
// 	             = 3 so f("cbd") < f("zaaaz").
//
// Example 2:
//
//	Input: queries = ["bbb","cc"], words = ["a","aa","aaa","aaaa"]
// 	Output: [1,2]
// 	Explanation: On the first query only f("bbb") < f("aaaa"). On
// 	             the second query both f("aaa") and f("aaaa") are
// 	             both > f("cc").
//
//
//
// Constraints:
//
//	1. 1 <= queries.length <= 2000
//	2. 1 <= words.length <= 2000
//	3. 1 <= queries[i].length, words[i].length <= 10
//	4. queries[i][j], words[i][j] are English lowercase letters.
//
// https://leetcode.com/problems/compare-strings-by-frequency-of-the-smallest-character

// Maximum number of English alphabets.
var maxAlphabets = 26

// Returns the count/frequency of smallest character in s.
// s is comprised only of lowercase English alphabets.
//
// E.g., returns 1 for s="really" as 'a' is the smallest and appears
// only once.
func smallCharCount(s string) int {
	small := 'z'
	freq := make([]int, maxAlphabets)
	for _, c := range s {
		if c < small {
			small = c
		}
		freq[int(c-'a')]++ // ASCII to integer for indexing
	}

	return freq[int(small-'a')]
}

func smallCharCounts(strings []string, c chan int) {
	for _, s := range strings {
		c <- smallCharCount(s)
	}
	close(c)
}

func numSmallerByFrequencyNaive(queries []string, words []string) []int {
	queryCount := make([]int, 0)
	queryChan := make(chan int, len(queries))
	go smallCharCounts(queries, queryChan)
	for i := range queryChan {
		queryCount = append(queryCount, i)
	}

	wordCount := make([]int, 0)
	wordChan := make(chan int, len(words))
	go smallCharCounts(words, wordChan)
	for i := range wordChan {
		wordCount = append(wordCount, i)
	}

	res := make([]int, len(queries))
	for i, qc := range queryCount {
		for _, wc := range wordCount {
			if qc < wc {
				res[i]++
			}
		}
	}

	return res
}

// https://leetcode.com/problems/compare-strings-by-frequency-of-the-smallest-character/discuss/366698/C++-Keep-track-of-count-of-frequencies-Beats-100
func numSmallerByFrequency(queries []string, words []string) []int {
	// Allocate word count for 12 items.
	// Constraint #3 guarantees max. words will be 10, but we use
	// 0-based index and use one additional slot for cumulative
	// frequency.
	wordCount := make([]int, 12)
	for _, word := range words {
		wordCount[smallCharCount(word)]++
	}

	// Compute cumulative word counts.
	// i.e., i'th word will have count of sum(wordCount[i:])
	for i := 9; i >= 0; i-- {
		wordCount[i] += wordCount[i+1]
	}

	res := make([]int, len(queries))
	for i, query := range queries {
		// If query's smallCount is 'c', then the number of
		// words that has smallCount > 'c' would be the `c+1`'th
		// wordCount since wordCount is now cumulative.
		res[i] = wordCount[smallCharCount(query)+1]
	}

	return res
}

func main() {
}
