package main

import (
	"github.com/prog-1/list"
)

// LessFunc returns true if a is smaller than b.
type LessFunc func(a, b int) bool

// InsertionSort rearranges the list to become sorted.
func InsertionSort(l *list.List, less LessFunc) {
}

// SortedInsert inserts a new element with a given value in a sorted list
// (while preserving a sorted order);
func SortedInsert(l *list.List, v int, less LessFunc) *list.Element {
	if l.Front() == nil { // no elements
		l.PushBack(v)
		return l.Back()
	}

	if less(v, l.Front().Value) { // must go first
		l.PushFront(v)
		return l.Front()
	}

	cur := l.Front() // cur will be the last element with value smaller or equal to v
	for cur.Next() != nil && less(cur.Next().Value, v) {
		cur = cur.Next()
	}

	return l.InsertAfter(v, cur)

}

func main() {
	var l list.List
	l.PushBack(0)
	l.PushBack(1)
	//InsertionSort(&l, func(a, b int) bool { return a < b })
	SortedInsert(&l, 2, func(a, b int) bool { return a < b })
	list.Print(&l)
}
