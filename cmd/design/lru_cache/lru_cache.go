package main

import "fmt"

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

type Node struct {
	key   int
	value int
	next  *Node
	prev  *Node
}

func NewLRUCache(capacity int) *LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head

	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node, capacity),
		head:     head,
		tail:     tail,
	}
}

func (lc *LRUCache) put(key, value int) {

}

func (lc *LRUCache) get(key int) int {
	return -1
}

func main() {

	lruCache := NewLRUCache(2)

	lruCache.put(1, 1)           // cache is {1=1}
	lruCache.put(2, 2)           // cache is {1=1, 2=2}
	fmt.Println(lruCache.get(1)) // return 1

	lruCache.put(3, 3)           // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	fmt.Println(lruCache.get(2)) // returns -1 (not found)
	lruCache.put(4, 4)           // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	fmt.Println(lruCache.get(1)) // return -1 (not found)
	fmt.Println(lruCache.get(3)) // return 3
	fmt.Println(lruCache.get(4)) // return 4
}
