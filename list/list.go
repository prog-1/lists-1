package list

import (
	"fmt"
)

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
func (l *List) PushFront(v int) {
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
func (l *List) PushBack(v int) {
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
func (l *List) Find(v int) *Element {
	for cur := l.head; cur != nil; cur = cur.next {
		if cur.Value == v {
			return cur
		}
	}
	return nil
}

// InsertAfter inserts a new element with a given value after
// a particular element.
func (l *List) InsertAfter(v int, prev *Element) *Element {
	e := NewElement(v)
	e.next = prev.next
	prev.next = e
	return e
}

// It is here because seemes like we can't access functions from *_test files
func ListToSlice(l *List) []int {
	s := make([]int, 0) // NOTE: not var s []int, because by default it gets initialized with nil, but I wanted to have an empty slice
	for cur := l.head; cur != nil; cur = cur.Next() {
		s = append(s, cur.Value)
	}
	return s
}

// Removes Element from List
func (l *List) Remove(el *Element) {
	if el == l.Front() { // is first
		l.head = l.Front().next
	} else {
		for prev := l.Front(); prev.Next() != nil; prev = prev.Next() {
			if prev.Next() == el {
				prev.next = el.next
				return
			}
		}
	}
}
