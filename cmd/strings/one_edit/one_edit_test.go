package main

import "testing"

func TestOneEdit(t *testing.T) {

	tests := []struct {
		name string
		s1   string
		s2   string
		exp  bool
	}{
		{
			name: "happy path",
			s1:   "hello",
			s2:   "helo",
			exp:  true,
		},
		{
			name: "s1 longer s2",
			s1:   "hello",
			s2:   "hel",
			exp:  false,
		},
		{
			name: "s1 is the same s2",
			s1:   "hello",
			s2:   "hello",
			exp:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := oneEdit(tt.s1, tt.s2)

			if res != tt.exp {
				t.Errorf("expected %v, but got %v", tt.exp, res)
			}
		})
	}
}
