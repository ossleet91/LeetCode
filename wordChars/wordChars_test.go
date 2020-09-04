package main

import (
	"testing"
)

func TestCountMap(t *testing.T) {
	// Create
	s := "hello"
	cm := newCountMap(s)
	if cm['h'] != 1 || cm['e'] != 1 || cm['l'] != 2 || cm['o'] != 1 {
		t.Errorf("incorrect counts after newCountMap for %s: cm=%#v\n",
			s, cm)
	}

	// Add
	w := "world"
	for _, r := range w {
		cm.Add(r)
	}
	if cm['w'] != 1 || cm['o'] != 2 || cm['r'] != 1 || cm['l'] != 3 || cm['d'] != 1 {
		t.Errorf("incorrect counts after adding %s: cm=%#v\n", w, cm)
	}

	// Delete
	cm.Del('o')
	if cm['w'] != 1 || cm['o'] != 1 || cm['r'] != 1 || cm['l'] != 3 || cm['d'] != 1 {
		t.Errorf("incorrect counts after deleting %c: cm=%#v\n", 'o', cm)
	}

	// Delete a rune not present in map. Should be no-op.
	cm.Del('m')
	if cm['w'] != 1 || cm['o'] != 1 || cm['r'] != 1 || cm['l'] != 3 || cm['d'] != 1 {
		t.Errorf("incorrect counts after deleting a rune not in map: cm=%#v\n",
			cm)
	}

	// DelFull
	cm.DelFull('l')
	if cm['w'] != 1 || cm['o'] != 1 || cm['r'] != 1 || cm['l'] != 0 || cm['d'] != 1 {
		t.Errorf("incorrect counts after fully deleting %c: cm=%#v\n", 'l', cm)
	}

	// Present
	if !cm.Present('w') {
		t.Errorf("map says a present value, '%c' is not present: cm=%#v\n", 'w', cm)
	}

	if cm.Present('z') {
		t.Errorf("map says a non-present value, '%c' is present: cm=%#v\n", 'z', cm)
	}
}

type TestCase struct {
	words []string
	chars string
	want  int
}

func TestCountCharacters(t *testing.T) {
	tests := []TestCase{
		{[]string{"cat", "bt", "hat", "tree"}, "atach", 6},
		{[]string{"hello", "world", "leetcode"}, "welldonehoneyr", 10},
	}

	for _, tt := range tests {
		got := countCharacters(tt.words, tt.chars)
		if got != tt.want {
			t.Errorf("failed for %#v: got=%d\n", tt, got)
		}
	}
}

func TestCountCharactersCM(t *testing.T) {
	tests := []TestCase{
		{[]string{"cat", "bt", "hat", "tree"}, "atach", 6},
		{[]string{"hello", "world", "leetcode"}, "welldonehoneyr", 10},
	}

	for _, tt := range tests {
		got := countCharactersCM(tt.words, tt.chars)
		if got != tt.want {
			t.Errorf("failed for %#v: got=%d\n", tt, got)
		}
	}
}
