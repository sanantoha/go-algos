package main

import "fmt"

func topKFrequent(words []string, k int) []string {
	return nil
}

/**
 * https://leetcode.com/explore/interview/card/apple/346/sorting-and-searching/3133/
 */
func main() {

	words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	fmt.Println(topKFrequent(words, 2))

	words1 := []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}
	fmt.Println(topKFrequent(words1, 4))
}
