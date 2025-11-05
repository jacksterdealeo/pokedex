package main

import (
	"testing"
)

func TestCases(t *testing.T) {
	var cases = []struct {
		input    string
		expected []string
	}{
		{
			input:    " ",
			expected: []string{},
		},
		{
			"red",
			[]string{"red"},
		},
		{
			"sMoThEr it With JAM ",
			[]string{"smother", "it", "with", "jam"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Fatalf("size of '%v' does not match '%v'.", actual, c.expected)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Fatalf("word '%v' does not match word '%v'", word, expectedWord)
			}
		}
	}
}
