package main

// Find words that can be formed by characters.
//
// You are given an array of strings words and a string chars.
//
// A string is good if it can be formed by characters from chars (each
// character can only be used once).
//
// Return the sum of lengths of all good strings in words.
//
// Example 1:
//	Input: words = ["cat","bt","hat","tree"], chars = "atach"
//	Output: 6
//	Explanation: The strings that can be formed are "cat" and "hat"
//		     so the answer is 3 + 3 = 6.
//
// Example 2:
//
//	Input: words = ["hello","world","leetcode"], chars = "welldonehoneyr"
//	Output: 10
//	Explanation: The strings that can be formed are "hello" and
//		     "world" so the answer is 5 + 5 = 10.
//
// Note:
//
//	1. 1 <= words.length <= 1000
//	2. 1 <= words[i].length, chars.length <= 100
//	3. All strings contain lowercase English letters only.
//
// https://leetcode.com/problems/find-words-that-can-be-formed-by-characters

// CountMap is a map of value to count.
//
// Useful to store how many times a value appears in a collection.
//
// E.g., CountMap of runes in "hello" will be {h:1, e:1, l:2, o:1}.
type CountMap map[rune]int

// Create a return a new CountMap for given string.
func newCountMap(s string) CountMap {
	cm := CountMap{}

	for _, r := range s {
		cm.Add(r)
	}

	return cm
}

// Add given rune to CountMap.
func (cm CountMap) Add(r rune) {
	cm[r]++
}

// Delete a given number from CountMap. If count is 1, delete the rune.
// Otherwise, decrement its count.
func (cm CountMap) Del(r rune) {
	if count, ok := cm[r]; ok {
		if count > 1 {
			cm[r]--
		} else {
			cm.DelFull(r)
		}
	}
}

// Delete a given rune from CountMap.
func (cm CountMap) DelFull(r rune) {
	delete(cm, r)
}

// Check if a given rune is present in CountMap.
func (cm CountMap) Present(r rune) bool {
	_, ok := cm[r]
	return ok
}

// Max. number of alphabets in English language.
const maxAlphabets = 26

func getCounts(s string) []int {
	// s can only comprise of lower-case English alphabets. See note
	// #3. So, allocating memory for just maxAlphabets counters is
	// enough to cover the string.
	counts := make([]int, maxAlphabets)
	for _, c := range s {
		counts[int(c-'a')]++ // ASCII math to find index
	}

	return counts
}

// Using simple array counter method.
func countCharacters(words []string, chars string) int {
	charCounts := getCounts(chars)
	count := 0

	for _, word := range words {
		wordCounts := getCounts(word)
		wordFromChars := true

		// Check if each character is present in CountMap. For
		// every character present, delete that character.
		// Stop once all characters are processed or if a
		// character is not present in CountMap.

		// 'word' can be made out of 'chars' if every letter in
		// word has same or lesser count of chars.
		for i := 0; i < maxAlphabets; i++ {
			if charCounts[i] < wordCounts[i] {
				wordFromChars = false
				break
			}
		}

		// Add the length of current word to result count only
		// if all characters are present in CountMap.
		if wordFromChars {
			count += len(word)
		}
	}

	return count
}

// Returns the length of word if word can be formed using letters in
// chars. Returns 0 otherwise. All returns are via channels.
func countIfWordFromChars(words []string, charCounts []int, c chan int) {
	for _, word := range words {
		wordCounts := getCounts(word)
		count := len(word)

		// 'word' can be made out of 'chars' if every letter in
		// word has same or lesser count of chars.
		for i := 0; i < maxAlphabets; i++ {
			if charCounts[i] < wordCounts[i] {
				count = 0
				break
			}
		}

		c <- count
	}

	close(c)
}

// Using simple array counter method and goroutines.
func countCharactersGoRoutine(words []string, chars string) int {
	charCounts := getCounts(chars)
	c := make(chan int, len(words))
	count := 0

	go countIfWordFromChars(words, charCounts, c)
	for i := range c {
		count += i
	}

	return count
}

// Using CountMap.
func countCharactersCM(words []string, chars string) int {
	count := 0

	for _, word := range words {
		cm := newCountMap(chars)
		wordFromString := true

		// Check if each character is present in CountMap. For
		// every character present, delete that character.
		// Stop once all characters are processed or if a
		// character is not present in CountMap.
		for _, r := range word {
			if cm.Present(r) {
				cm.Del(r)
			} else {
				wordFromString = false
				break
			}
		}

		// Add the length of current word to result count only
		// if all characters are present in CountMap.
		if wordFromString {
			count += len(word)
		}
	}

	return count
}

func main() {
}
