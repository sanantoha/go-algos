package main

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

func TestIsMatch(t *testing.T) {

	tests := []struct {
		name string
		s    string
		p    string
		exp  bool
	}{
		{
			name: "s=aa, p=a",
			s:    "aa",
			p:    "a",
			exp:  false,
		},
		{
			name: "s=aa, p=a*",
			s:    "aa",
			p:    "a*",
			exp:  true,
		},
		{
			name: "s=abcde, p=.*",
			s:    "abcde",
			p:    ".*",
			exp:  true,
		},
		{
			name: "s=abcde, p=.*de",
			s:    "abcde",
			p:    ".*de",
			exp:  true,
		},
		{
			name: "s=abcde, p=.*dk",
			s:    "abcde",
			p:    ".*dk",
			exp:  false,
		},
	}

	funcs := []func(string, string) bool{
		isMatch,
		isMatchIter,
	}

	for _, tt := range tests {
		for _, fn := range funcs {
			pc := reflect.ValueOf(fn).Pointer()
			fun := runtime.FuncForPC(pc)

			testName := fmt.Sprintf("%s %s", fun.Name(), tt.name)

			t.Run(testName, func(t *testing.T) {
				res := fn(tt.s, tt.p)

				if tt.exp != res {
					t.Errorf("expected %v, but got %v", tt.exp, res)
				}
			})
		}
	}
}
