package main

import (
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func TestFindPermutation(t *testing.T) {
	for _, fn := range []func(string, string) bool{findPermutation, findPermutation1} {
		testTable := []struct {
			name string
			s1   string
			s2   string
			exp  bool
		}{
			{
				name: "happy path for " + getFunctionName(fn),
				s1:   "abc",
				s2:   "hdflebacworld",
				exp:  true,
			},
			{
				name: "happy path1 " + getFunctionName(fn),
				s1:   "abbc",
				s2:   "hbbcadflebdworld",
				exp:  true,
			},
		}

		for _, subtest := range testTable {
			t.Run(subtest.name, func(t *testing.T) {
				res := fn(subtest.s1, subtest.s2)

				if subtest.exp != res {
					t.Errorf("expected %v, got %v", subtest.exp, res)
				}
			})
		}
	}
}

func getFunctionName(i interface{}) string {
	name := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	arr := strings.Split(name, ".")
	return arr[len(arr)-1]
}

//
//func TestFindPermutation1(t *testing.T) {
//	testTable := []struct {
//		name string
//		s1   string
//		s2   string
//		exp  bool
//	}{
//		{
//			name: "happy path",
//			s1:   "abc",
//			s2:   "hdflebacworld",
//			exp:  true,
//		},
//		{
//			name: "happy path1",
//			s1:   "abbc",
//			s2:   "hbbcadflebdworld",
//			exp:  true,
//		},
//	}
//
//	for _, subtest := range testTable {
//		t.Run(subtest.name, func(t *testing.T) {
//			res := findPermutation1(subtest.s1, subtest.s2)
//
//			if subtest.exp != res {
//				t.Errorf("expected %v, got %v", subtest.exp, res)
//			}
//		})
//	}
//}
