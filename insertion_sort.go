package main

import (
	"github.com/prog-1/list"
)

// LessFunc returns true if a is smaller than b.
type LessFunc func(a, b int) bool

// InsertionSort rearranges the list to become sorted.
func InsertionSort(l *list.List, less LessFunc) {
	var sorted *list.List = new(list.List)
	for cur := (*l).Front(); cur != nil; cur = cur.Next() {
		SortedInsert(sorted, cur.Value, func(a, b int) bool { return a < b })
	}

}

// SortedInsert inserts a new element with a given value in a sorted list
// (while preserving a sorted order);
func SortedInsert(l *list.List, v int, less LessFunc) *list.Element {
	cur := (*l).Front() // cur will be the last element with value smaller than v or equal to v
	for cur != nil && cur.Next() != nil && cur.Next().Value < v {
		cur = cur.Next()
	}
	switch cur {
	case nil:
		(*l).PushBack(v)
		return l.Back()
	case l.Front():
		if l.Front().Value > v {
			l.PushFront(v)
			return l.Front()
		} else {
			return (*l).InsertAfter(v, cur)
		}
	default:
		return (*l).InsertAfter(v, cur)
	}
}

func main() {
	var l list.List
	l.PushFront(4)
	l.PushFront(2)
	l.PushFront(1)
	l.PushFront(3)
	InsertionSort(&l, func(a, b int) bool { return a < b })
	list.Print(&l)
}
