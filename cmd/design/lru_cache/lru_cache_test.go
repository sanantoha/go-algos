package main

import (
	"fmt"
	"testing"
)

func TestLruCache(t *testing.T) {

	lruCache := NewLRUCache(2)

	lruCache.put(1, 1) // cache is {1=1}

	expCache := map[int]*Node{1: {key: 1, value: 1}}
	if !CompareMaps(lruCache.cache, expCache) {
		t.Errorf("expected %v, but got %v", expCache, lruCache.cache)
	}

	lruCache.put(2, 2) // cache is {1=1, 2=2}

	expCache = map[int]*Node{1: {key: 1, value: 1}, 2: {key: 2, value: 2}}
	if !CompareMaps(lruCache.cache, expCache) {
		t.Errorf("expected %v, but got %v", expCache, lruCache.cache)
	}

	res := lruCache.get(1)
	fmt.Println(res) // return 1
	if res != 1 {
		t.Errorf("expected 1, but got %v", res)
	}

	lruCache.put(3, 3) // LRU key was 2, evicts key 2, cache is {1=1, 3=3}

	expCache = map[int]*Node{1: {key: 1, value: 1}, 3: {key: 3, value: 3}}
	if !CompareMaps(lruCache.cache, expCache) {
		t.Errorf("expected %v, but got %v", expCache, lruCache.cache)
	}

	res = lruCache.get(2)
	fmt.Println(res) // return -1
	if res != -1 {
		t.Errorf("expected -1, but got %v", res)
	}

	lruCache.put(4, 4) // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	expCache = map[int]*Node{4: {key: 4, value: 4}, 3: {key: 3, value: 3}}
	if !CompareMaps(lruCache.cache, expCache) {
		t.Errorf("expected %v, but got %v", expCache, lruCache.cache)
	}

	res = lruCache.get(1) // return -1 (not found)
	fmt.Println(res)      // return -1
	if res != -1 {
		t.Errorf("expected -1, but got %v", res)
	}

	res = lruCache.get(3)
	fmt.Println(res) // return 3
	if res != 3 {
		t.Errorf("expected 3, but got %v", res)
	}
	res = lruCache.get(4)
	fmt.Println(res) // return 4
	if res != 4 {
		t.Errorf("expected 4, but got %v", res)
	}
}

// CompareMaps compares two maps with int keys and *Node values
func CompareMaps(map1, map2 map[int]*Node) bool {
	// Check if both maps are nil
	if map1 == nil && map2 == nil {
		return true
	}
	// Check if one map is nil and the other is not
	if map1 == nil || map2 == nil {
		return false
	}
	// Check if lengths are different
	if len(map1) != len(map2) {
		return false
	}

	// Iterate over map1 and compare with map2
	for key, node1 := range map1 {
		node2, exists := map2[key]
		// Check if key exists in map2
		if !exists {
			return false
		}
		// Check for nil nodes
		if node1 == nil && node2 == nil {
			continue
		}
		if node1 == nil || node2 == nil {
			return false
		}
		// Compare Node fields
		if node1.key != node2.key || node1.value != node2.value {
			return false
		}
	}

	return true
}
