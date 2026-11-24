package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " hello    world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HellO WORLD",
			expected: []string{"hello", "world"},
		},
		{
			input:    "more Than two WORDS",
			expected: []string{"more", "than", "two", "words"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// check length
		if len(actual) != len(c.expected) {
			t.Error("different lengths")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				err := fmt.Errorf("'%s' != '%s'", word, expectedWord)
				fmt.Printf("err: %v\n", err)
			}

		}
	}

}
