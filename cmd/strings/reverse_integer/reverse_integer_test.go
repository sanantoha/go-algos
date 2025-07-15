package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name  string
		input int
		exp   int
	}{
		{
			name:  "123",
			input: 123,
			exp:   321,
		},
		{
			name:  "0",
			input: 0,
			exp:   0,
		},
		{
			name:  "-123",
			input: -123,
			exp:   -321,
		},
		{
			name:  "120",
			input: 120,
			exp:   21,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := reverse(tt.input)

			if res != tt.exp {
				t.Errorf("expected %v, but got %v", tt.exp, res)
			}
		})
	}
}
