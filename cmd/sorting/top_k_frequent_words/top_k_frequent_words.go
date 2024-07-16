package main

import (
	"container/heap"
	"fmt"
	"slices"
	"strings"
)

// O(n * log(k)) time | O(n) space
func topKFrequent(words []string, k int) []string {
	if words == nil || len(words) == 0 || k < 0 {
		return nil
	}
	mp := make(map[string]*Word, 0)

	for _, w := range words {
		word, ok := mp[w]
		if ok {
			word.cnt++
		} else {
			mp[w] = &Word{
				w:   w,
				cnt: 1,
			}
		}
	}

	h := &WordHeap{}
	i := 0
	for _, word := range mp {
		heap.Push(h, *word)

		if i >= k {
			heap.Pop(h)
		}
		i++
	}

	res := make([]string, 0)

	for h.Len() > 0 {
		word := heap.Pop(h).(Word)
		res = append(res, word.w)
	}

	slices.Reverse(res)
	return res
}

// Heap implementation.
type Word struct {
	w   string
	cnt int
}

type WordHeap []Word

func (h WordHeap) Len() int {
	return len(h)
}

func (h WordHeap) Less(i, j int) bool {
	if h[i].cnt == h[j].cnt {
		return strings.Compare(h[i].w, h[j].w) > 0
	}
	return h[i].cnt < h[j].cnt
}

func (h WordHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *WordHeap) Push(x interface{}) {
	*h = append(*h, x.(Word))
}

func (h *WordHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

/**
 * https://leetcode.com/explore/interview/card/apple/346/sorting-and-searching/3133/
 */
func main() {

	words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	fmt.Println(topKFrequent(words, 2)) // {"i", "love"}

	words1 := []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}
	fmt.Println(topKFrequent(words1, 4)) // {"the","is","sunny","day"}
}
