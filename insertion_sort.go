package main

import (
	"fmt"

	"github.com/prog-1/list"
)

// LessFunc returns true if a is smaller than b.
type LessFunc func(a, b int) bool

// InsertionSort rearranges the list to become sorted.
func InsertionSort(l *list.List, less LessFunc) {
	temp := l.Front()
	for temp.Next() != nil {

	}
}

// SortedInsert inserts a new element with a given value in a sorted list
// (while preserving a sorted order);
func SortedInsert(l *list.List, v int, less LessFunc) *list.Element {
	// TODO
	return nil
}

func main() {
	var l list.List
	l.PushBack(1)
	l.PushBack(3)
	l.PushBack(2)
	l.PushFront(4)
	for n := l.Front(); n != nil; n = n.Next() {
		fmt.Println(n.Value)
	}
}
