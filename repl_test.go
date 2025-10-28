package main

import "testing"

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "my.name.is.jake",
			expected: []string{"my.name.is.jake"},
		},
		{
			input:    "LOWERCASE THIS",
			expected: []string{"lowercase", "this"},
		},
		{
			input:    "t h i s I S",
			expected: []string{"t", "h", "i", "s", "i", "s"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("actual len: %d is not equal expected len: %d", len(actual), len(c.expected))
		}
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("actual char: %v does not match expected char: %v", word, expectedWord)
			}
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
	}
}
