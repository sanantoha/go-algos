package main

import (
	"fmt"
	"testing"
)

func TestStaircaseTraversal(t *testing.T) {
	testTable := []struct {
		name     string
		height   int
		maxSteps int
		expRes   int
	}{
		{
			name:     "height=4, maxSteps=2",
			height:   4,
			maxSteps: 2,
			expRes:   5,
		},
		{
			name:     "height=4, maxSteps=3",
			height:   4,
			maxSteps: 3,
			expRes:   7,
		},
	}

	for _, test := range testTable {
		testName := fmt.Sprintf("staircaseTraversalRec %s", test.name)
		t.Run(testName, func(t *testing.T) {
			res := staircaseTraversalRec(test.height, test.maxSteps)

			if test.expRes != res {
				t.Errorf("expected %v, but got %v", test.expRes, res)
			}
		})

		testName = fmt.Sprintf("staircaseTraversalRecMemoization %s", test.name)
		t.Run(testName, func(t *testing.T) {
			res := staircaseTraversalRecMemoization(test.height, test.maxSteps)

			if test.expRes != res {
				t.Errorf("expected %v, but got %v", test.expRes, res)
			}
		})

		testName = fmt.Sprintf("staircaseTraversalIter %s", test.name)
		t.Run(testName, func(t *testing.T) {
			res := staircaseTraversalIter(test.height, test.maxSteps)

			if test.expRes != res {
				t.Errorf("expected %v, but got %v", test.expRes, res)
			}
		})

		testName = fmt.Sprintf("staircaseTraversalSlidingWindow %s", test.name)
		t.Run(testName, func(t *testing.T) {
			res := staircaseTraversalSlidingWindow(test.height, test.maxSteps)

			if test.expRes != res {
				t.Errorf("expected %v, but got %v", test.expRes, res)
			}
		})
	}

}
