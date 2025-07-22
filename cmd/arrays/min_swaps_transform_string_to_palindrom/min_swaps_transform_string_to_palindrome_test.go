package main

import (
	"testing"
)

func TestMinSwapsRequired(t *testing.T) {

	tests := []struct {
		name  string
		input string
		exp   int
	}{
		{
			name:  "happy path",
			input: "0100101",
			exp:   2,
		},
		{
			name:  "impossible to swap",
			input: "1110",
			exp:   -1,
		},
		{
			name:  "one swap only",
			input: "11101",
			exp:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := minSwapsRequired(tt.input)

			if res != tt.exp {
				t.Errorf("expected %v, but got %v", tt.exp, res)
			}
		})
	}
}
