package main

import (
	"github.com/prog-1/list"
)

// LessFunc returns true if a is smaller than b.
type LessFunc func(a, b int) bool

// InsertionSort rearranges the list to become sorted.
func InsertionSort(l *list.List) *list.List {
	sl := &list.List{} //sorted list
	for el := l.Front(); el != nil; el = el.Next() {
		SortedInsert(sl, el.Value)
	}
	return sl
}

// SortedInsert inserts a new element with a given value in a sorted list
// (while preserving a sorted order);
func SortedInsert(l *list.List, v int) *list.Element {

	//if given val is smallest
	//or slice is empty
	if l.Len() == 0 || v <= l.Front().Value {
		l.PushFront(v)
		return l.Front().Next()
	}
	// else if v > l.Back().Value {
	// 	//if given val is biggest
	// 	l.PushBack(v)
	// 	return l.Back()
	// }

	el := l.Front()
	for el.Next() != nil && el.Next().Value < v {
		el = el.Next()
	}
	return l.InsertAfter(v, el)
}

// func main() {
// 	l := &list.List{}
// 	l.PushFront(2)
// 	l.PushFront(3)
// 	l.PushFront(5)
// 	l.PushFront(1)
// 	l.PushFront(4)
// 	//list.Print(l)
// 	l = InsertionSort(l)
// 	list.Print(l)
// }
