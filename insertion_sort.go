package main

import (
	"fmt"
)

type T int

// Element is an element of a singly linked list.
type Element struct {
	// The value stored with this element.
	Value T
	next  *Element
}

// NewElement constructs and returns a new Element with a given value.
func NewElement(v T) *Element {
	return &Element{Value: v}
}

// Next returns the next list element or nil.
func (e *Element) Next() *Element {
	return e.next
}

// List represents a singly linked list.
// https://en.wikipedia.org/wiki/Linked_list
type List struct {
	head, tail *Element
	len        uint
}

// Front returns the first element of list or nil if the list is empty.
func (l *List) Front() *Element {
	return l.head
}

// Back returns the last element of list or nil if the list is empty.
func (l *List) Back() *Element {
	return l.tail
}

// Print prints all nodes of the linked list.
func Print(l *List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
		if e.Next() != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}

// PushFront inserts a new element at the front of list.
func (l *List) PushFront(v T) {
	e := NewElement(v)
	if l.len == 0 {
		l.head, l.tail = e, e
	} else {
		e.next = l.head
		l.head = e
	}
	l.len++
}

// PushBack inserts a new element at the back of list.
func (l *List) PushBack(v T) {
	e := NewElement(v)
	if l.len == 0 {
		l.head, l.tail = e, e
	} else {
		l.tail.next = e
		l.tail = e
	}
	l.len++
}

// Find returns an element with a given value from the list or nil,
// if the value is not found.
func (l *List) Find(v T) *Element {
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == v {
			return e
		}
	}
	return nil
}

type LessFunc func(a, b int) bool

// InsertAfter inserts a new element with a given value after
// a particular element.
func (l *List) InsertAfter(v T, prev *Element) *Element {
	e := NewElement(v)
	e.next = prev.next
	prev.next = e
	l.len++
	return e
}

// SortedInsert inserts a new element with a given value in a sorted list
// (while preserving a sorted order);
func (l *List) SortedInsert(v T) *Element {
	x := l.Front()
	if l.len == 0 {
		l.PushFront(v)
		return l.Front()
	}
	for x != nil && x.next.Value < v {
		x = x.next
	}
	return l.InsertAfter(v, x)
}

// InsertionSort sorts a list using an insertion sort.
func InsertionSort(l *List) *List {
	lol := &List{}
	for e := l.Front(); e != nil; e = e.Next() {
		lol.SortedInsert(e.Value)
	}
	return lol
}

func main() {
	// l1 := &List{}
	// l1.PushBack(101)
	// l1.PushFront(12)
	// l1.PushFront(20)
	// l1.PushFront(25)
	// l1.PushBack(100)
	// Print(l1) // 25 -> 20 -> 12
	// l1.InsertAfter(50, l1.Front())
	// // 25 -> 50 -> 20 -> 12
	l := &List{}
	l.PushBack(1)
	l.PushBack(3)
	l.PushBack(4)
	l.PushFront(0)
	// Print(l)
	// fmt.Println(l.Find(5))
	// fmt.Println(l.InsertAfter(100, l.head))
	// Print(l)
	// fmt.Println(l.SortedInsert(2))
	// Print(l)
	Print(InsertionSort(l))

	// for n := l.Front(); n != nil; n = n.Next() {
	// 	fmt.Println(n.Value)
	// }
}
