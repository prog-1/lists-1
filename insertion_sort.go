package main

import (
	"github.com/prog-1/list"
)

// LessFunc returns true if a is smaller than b.
type LessFunc func(a, b int) bool

// InsertionSort rearranges the list to become sorted.
func InsertionSort(l *list.List, less LessFunc) {
	if l.Front() == l.Back() { // 1 or less elements
		return
	}

	for i := l.Front().Next(); i != nil; i = i.Next() {
		if less(i.Value, l.Front().Value) { // must go first
			l.PushFront(i.Value)
			l.Remove(i)
		} else { // finding el after which we need to insert i and moving it there
			for j := l.Front().Next(); j.Next() != i; j = j.Next() {
				if less(i.Value, j.Next().Value) {
					l.InsertAfter(i.Value, j)
					l.Remove(i)
					break
				}
			}
		}
		/*l.Remove(i)*/ // For some reason when it is here it does not execute properly
	}
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
	l.PushBack(1)
	l.PushBack(0)
	l.PushBack(3)
	l.PushBack(2)
	InsertionSort(&l, func(a, b int) bool { return a < b })
	//SortedInsert(&l, 2, func(a, b int) bool { return a < b })
	//l.Remove(l.Find(2))
	list.Print(&l)
}
