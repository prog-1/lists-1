package main

import (
	"fmt"

	"github.com/prog-1/list"
)

// LessFunc returns true if a is smaller than b.
type LessFunc func(a, b int) bool

// InsertionSort rearranges the list to become sorted.
func InsertionSort(l *list.List, less LessFunc) {
	var tmp list.List
	for n := l.Front(); n != nil; n = n.Next() {
		SortedInsert(&tmp, n.Value, less)
		for n := tmp.Front(); n != nil; n = n.Next() {
			fmt.Println(n.Value)
		}
	}
	// for n := tmp.Front(); n != nil; n = n.Next() {
	// 	fmt.Println(n.Value)
	// }
	l.Copy(tmp)
}

// SortedInsert inserts a new element with a given value in a sorted list
// (while preserving a sorted order);
func SortedInsert(l *list.List, v int, less LessFunc) *list.Element {
	for n := l.Front(); n != nil; n = n.Next() {
		if n.Next() == nil || !less(n.Next().Value, v) {
			return l.InsertAfter(v, n)
		}
	}
	l.PushBack(v)
	return nil
}

func main() {
	var l list.List
	l.PushBack(4)
	l.PushBack(2)
	l.PushBack(1)
	for n := l.Front(); n != nil; n = n.Next() {
		fmt.Println(n.Value)
	}
	less := func(a, b int) bool {
		return a < b
	}
	InsertionSort(&l, less)
	for n := l.Front(); n != nil; n = n.Next() {
		fmt.Println(n.Value)
	}
}
