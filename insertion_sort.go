package main

import (
	"github.com/prog-1/list"
)

// LessFunc returns true if a is smaller than b.
type LessFunc func(a, b int) bool

// InsertionSort rearranges the list to become sorted.
func InsertionSort(l *list.List, less LessFunc) {
	for n := l.Front().Next(); n != nil; n = n.Next() {
		for k := n; k != l.Front() && k.Previous().Value > k.Value; k = k.Previous() {
			l.SwapElements(k.Previous().Value, k.Value)
		}
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

	l.PushBack(123)
	//l.InsertAfter(l.RemoveAfter(nil), nil)
	//moveNextElement(&l, nil, nil)
	list.Print(&l)

	for i := 1; i < 4; i++ {
		l.PushBack(i)
		l.PushFront(i + 3)
	}
	list.Print(&l)

	less := func(a, b int) bool { return a > b }

	InsertionSort(&l, less)
	list.Print(&l)

	for i := 0; i < 10; i++ {
		SortedInsert(&l, i, less)
	}
	list.Print(&l)
}
