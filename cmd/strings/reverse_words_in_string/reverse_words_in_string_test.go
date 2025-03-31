package main

import "testing"

func TestReverse(t *testing.T) {
	testTable := []struct {
		name      string
		input     string
		expOutput string
	}{
		{
			name:      "happy path",
			input:     "the sky is blue",
			expOutput: "blue is sky the",
		},
		{
			name:      "string with spaces",
			input:     " a good   example     ",
			expOutput: "     example   good a ",
		},
		{
			name:      "more spaces between words",
			input:     "this      string     has a     lot of   whitespace",
			expOutput: "whitespace   of lot     a has     string      this",
		},
		{
			name:      "empty string",
			input:     "",
			expOutput: "",
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := reverse(subtest.input)

			if subtest.expOutput != res {
				t.Errorf("expected (%v), got (%v)", subtest.expOutput, res)
			}
		})
	}
}
