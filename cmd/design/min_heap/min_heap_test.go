package main

import (
	"reflect"
	"testing"
)

func TestBuildHeap(t *testing.T) {
	expected := []int{2, 4, 3, 5, 8, 6, 4, 9, 7}
	heap := []int{9, 4, 6, 7, 8, 3, 4, 5, 2}

	res := buildHeap(heap)

	if !reflect.DeepEqual(expected, res) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestSiftDown(t *testing.T) {
	heap := []int{2, 4, 3, 5, 8, 6, 4, 9, 7}
	heap = append([]int{10}, heap...)
	expected := []int{2, 3, 4, 4, 5, 8, 6, 10, 9, 7}

	siftDown(0, len(heap)-1, heap)

	if !reflect.DeepEqual(expected, heap) {
		t.Errorf("expected %v, got %v", expected, heap)
	}
}

func TestSiftUp(t *testing.T) {
	heap := []int{2, 4, 3, 5, 8, 6, 4, 9, 7}
	heap = append(heap, 1)
	expected := []int{1, 2, 3, 5, 4, 6, 4, 9, 7, 8}

	siftUp(len(heap)-1, heap)

	if !reflect.DeepEqual(expected, heap) {
		t.Errorf("expected %v, got %v", expected, heap)
	}
}

func TestNewMinHeap(t *testing.T) {
	arr := []int{2, 1, 0}
	expected := []int{0, 1, 2}

	minHeap := NewMinHeap(arr)

	if !reflect.DeepEqual(expected, minHeap.heap) {
		t.Errorf("expected %v, got %v", expected, minHeap.heap)
	}
}

func TestInsert(t *testing.T) {
	arr := []int{2, 3, 4}
	expected := []int{1, 2, 4, 3}

	minHeap := NewMinHeap(arr)

	minHeap.Insert(1)

	if !reflect.DeepEqual(expected, minHeap.heap) {
		t.Errorf("expected %v, got %v", expected, minHeap.heap)
	}
}

func TestPeek(t *testing.T) {
	arr := []int{2, 3, 4}
	exp := 2

	minHeap := NewMinHeap(arr)

	peek := minHeap.Peek()

	if exp != peek {
		t.Errorf("expected %v, but got %v", exp, peek)
	}
}

func TestRemove(t *testing.T) {
	arr := []int{2, 3, 4}
	exp := 2
	expHeap := []int{3, 4}

	minHeap := NewMinHeap(arr)

	v := minHeap.Remove()

	if exp != v {
		t.Errorf("expected %v, but got %v", exp, v)
	}

	heap := minHeap.heap

	if !reflect.DeepEqual(expHeap, heap) {
		t.Errorf("expected %v, but got %v", exp, v)
	}
}

func TestIsEmpty(t *testing.T) {
	arr := []int{1, 2, 3}

	minHeap := NewMinHeap(arr)

	empty := minHeap.IsEmpty()

	if empty {
		t.Errorf("expected %v, but got %v", false, empty)
	}
}
