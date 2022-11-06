package main

import (
	"github.com/prog-1/list"
)

// LessFunc returns true if a is smaller than b.
type LessFunc func(a, b int) bool

// InsertionSort rearranges the list to become sorted.
func InsertionSort(l *list.List, less LessFunc) {
	// var n = l.Len()
	// e := l.Front()
	// for i := n - 1; i > 0; i-- {
	// 	for c := uint(0); i > c; c++ {
	// 		e = e.Next()
	// 	}
	// 	for j := i; j < n || e.Next() != nil; j++ {
	// 		if less(e.Next().Value, e.Value) {
	// 			e.Next().Value, e.Value = e.Value, e.Next().Value
	// 		}
	// 	}
	// 	e = l.Front()
	// }

	// 	if !less(e.Value, e.Next().Value) {
	// 		var tmp = l.Front()
	// 		var swapped bool
	// 		for j = 0; j < i; j++ {
	// 			if !less(e.Value, tmp.Value) && !swapped {
	// 				swapped = true
	// 				l.InsertAfter(e.Value, tmp)
	// 			}
	// 			tmp = tmp.Next()

	// 		}
	// 		if !swapped {

	// 		}

	// 	}
	// 	e = e.Next()
	// 	j = j - 1
	// }
}

// SortedInsert inserts a new element with a given value in a sorted list
// (while preserving a sorted order);
func SortedInsert(l *list.List, v int, less LessFunc) *list.Element {
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Next() == nil || less(v, e.Next().Value) {
			return l.InsertAfter(v, e)
		}
	}
	l.PushBack(v)
	return nil
}

func main() {
	var l list.List
	l.PushBack(1)
	l.PushBack(4)
	l.PushBack(3)
	l.PushBack(11)
	l.PushFront(4)
	InsertionSort(&l, func(a, b int) bool {
		return a < b
	})
	list.Print(&l)
}
