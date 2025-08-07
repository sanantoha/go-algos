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

func (n Node) String() string {
	return fmt.Sprintf("Node{key=%v, value=%v}", n.key, n.value)
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

// O(1) time | O(1) space
func (lc *LRUCache) put(key, value int) {
	node, ok := lc.cache[key]
	if ok {
		node.value = value
		remove(node)
		lc.add(node)
	} else {
		node = &Node{
			key:   key,
			value: value,
		}

		if lc.capacity == len(lc.cache) {
			delete(lc.cache, lc.tail.prev.key)
			remove(lc.tail.prev)
		}
		lc.add(node)
		lc.cache[key] = node
	}
}

// O(1) time | O(1) space
func (lc *LRUCache) get(key int) int {
	node, ok := lc.cache[key]
	if ok {
		remove(node)
		lc.add(node)
		return node.value
	} else {
		return -1
	}
}

func (lc *LRUCache) add(node *Node) {
	headNext := lc.head.next

	node.prev = lc.head
	node.next = headNext

	lc.head.next = node
	headNext.prev = node
}

func remove(node *Node) {
	next := node.next
	prev := node.prev

	prev.next = next
	next.prev = prev
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
