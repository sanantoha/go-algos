package main

import (
	"fmt"
	"reflect"
	"sort"
)

// O(n ^ 2) time | O(n ^ 2) space
func stableInternships(interns [][]int, teams [][]int) [][]int {
	chosenIntern := make(map[int]int, 0)
	freeInterns := make([]int, len(interns))
	for i, _ := range interns {
		freeInterns[i] = i
	}
	currentInternChoice := make([]int, len(interns))

	teamMaps := make([]map[int]int, len(teams))
	for i, team := range teams {
		rank := make(map[int]int)
		for j, v := range team {
			rank[v] = j
		}
		teamMaps[i] = rank
	}

	for len(freeInterns) != 0 {
		topIdx := len(freeInterns) - 1
		currentIntern := freeInterns[topIdx]
		freeInterns = freeInterns[:topIdx]

		intern := interns[currentIntern]
		teamPreference := intern[currentInternChoice[currentIntern]]
		currentInternChoice[currentIntern]++

		prevIntern, ok := chosenIntern[teamPreference]
		if !ok {
			chosenIntern[teamPreference] = currentIntern
			continue
		}

		prevInternRank := teamMaps[teamPreference][prevIntern]
		currInternRank := teamMaps[teamPreference][currentIntern]

		if currInternRank < prevInternRank {
			chosenIntern[teamPreference] = currentIntern
			freeInterns = append(freeInterns, prevIntern)
		} else {
			freeInterns = append(freeInterns, currentIntern)
		}

	}

	res := make([][]int, 0)
	for team, intern := range chosenIntern {
		res = append(res, []int{intern, team})
	}

	return res
}

func main() {
	interns := [][]int{
		{0, 1, 2},
		{0, 2, 1},
		{1, 2, 0},
	}

	teams := [][]int{
		{2, 1, 0},
		{0, 1, 2},
		{0, 1, 2},
	}

	expected := [][]int{
		{0, 1},
		{1, 0},
		{2, 2},
	}

	result := stableInternships(interns, teams)
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})

	fmt.Println(result)
	fmt.Println(reflect.DeepEqual(expected, result))
}
