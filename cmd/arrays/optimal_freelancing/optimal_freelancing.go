package main

import (
	"fmt"
	"math"
	"sort"
)

// O(n * log(n)) time | O(1) space
func optimalFreelancing(jobs []struct {
	deadline int
	payment  int
}) int {

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].payment > jobs[j].payment
	})

	lengthOfWeek := 7
	timelines := make([]bool, lengthOfWeek)

	maxProfit := 0

	for _, job := range jobs {
		maxTime := int(math.Min(float64(job.deadline), float64(lengthOfWeek)))
		for time := maxTime - 1; time >= 0; time-- {
			if !timelines[time] {
				timelines[time] = true
				maxProfit += job.payment
				break
			}
		}
	}

	return maxProfit
}

func main() {

	jobs := []struct {
		deadline int
		payment  int
	}{
		{
			deadline: 2,
			payment:  2,
		},
		{
			deadline: 4,
			payment:  3,
		},
		{
			deadline: 5,
			payment:  1,
		},
		{
			deadline: 7,
			payment:  2,
		},
		{
			deadline: 3,
			payment:  1,
		},
		{
			deadline: 3,
			payment:  2,
		},
		{
			deadline: 1,
			payment:  3,
		},
	}

	fmt.Println(optimalFreelancing(jobs)) // 13
}
