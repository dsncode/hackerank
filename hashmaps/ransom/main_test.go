package main

import "testing"

type TestCases struct {
	magazine []string
	note     []string
	expected bool
}

func TestNotes(t *testing.T) {

	testcases := []TestCases{
		{
			magazine: []string{"give", "me", "one", "grand", "today", "night"},
			note:     []string{"give", "one", "grand", "today"},
			expected: true,
		},
		{
			magazine: []string{"two", "times", "three", "is", "not", "four"},
			note:     []string{"two", "times", "two", "is", "four"},
			expected: false, // 'two' only occurs once in the magazine.

		},
		{
			magazine: []string{"ive", "got", "a", "lovely", "bunch", "of", "coconuts"},
			note:     []string{"ive", "got", "some", "coconuts"},
			expected: false, // Harold's magazine is missing the word.
		},
	}

	for _, tcase := range testcases {
		if checkWords(tcase.magazine, tcase.note) != tcase.expected {
			t.Fatal("wrong expectation")
		}
	}

}
