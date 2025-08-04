package main

import (
	"reflect"
	"runtime"
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	str1 := "ZXVVYZW"
	str2 := "XKYKZPW"
	expected := []rune{'X', 'Y', 'Z', 'W'}

	funcs := []func(string, string) []rune{
		longestCommonSubsequence,
		longestCommonSubsequence1,
	}

	for _, fn := range funcs {
		pc := reflect.ValueOf(fn).Pointer()
		fun := runtime.FuncForPC(pc)

		t.Run(fun.Name(), func(t *testing.T) {
			res := fn(str1, str2) // [88 89 90 87]

			if !reflect.DeepEqual(expected, res) {
				t.Errorf("expected %v, but got %v", expected, res)
			}
		})
	}

}
