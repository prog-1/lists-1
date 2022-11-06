package main

import (
	"fmt"

	"github.com/prog-1/list"
)

// LessFunc returns true if a is smaller than b.
type LessFunc func(a, b int) bool

// InsertionSort rearranges the list to become sorted.
func InsertionSort(l *list.List, less LessFunc) {
	head := l.Front()
	if head == nil {
		fmt.Println("Empty l.l")
	} else {
		for head != nil {
			j := head.Next()
			for j != nil {
				if less(j.Value, head.Value) {
					temp := head.Value
					head.Value = j.Value
					j.Value = temp
				}
				j = j.Next()
			}
			head = head.Next()
		}
	}
}

// SortedInsert inserts a new element with a given value in a sorted list
// (while preserving a sorted order);
func SortedInsert(l *list.List, v int, less LessFunc) *list.Element {
	head := l.Front()
	if head == nil {
		l.PushFront(v)
		return l.Front()
	}
	for head != nil {
		if !less(head.Next().Value, v) {
			head = head.Next()
		}
	}
	return l.InsertAfter(v, head)
}

func main() {
	var l list.List
	var less LessFunc
	l.PushBack(1)
	l.PushBack(3)
	l.PushBack(2)
	l.PushFront(4)
	for n := l.Front(); n != nil; n = n.Next() {
		fmt.Println(n.Value)
	}
	InsertionSort(&l, less)
	for n := l.Front(); n != nil; n = n.Next() {
		fmt.Println(n.Value)
	}
}
