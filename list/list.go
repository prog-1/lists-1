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

func (l *List) Len() uint {
	return l.len
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
		l.tail.next = e //changing link of current last element to new element
		l.tail = e      //changing last element to new one
	}
	l.len++
}

// Find returns an element with a given value from the list or nil,
// if the value is not found.
func (l *List) Find(v int) *Element {
	el := l.head
	for el != nil && el.Value != v {
		el = el.next
	}
	return el
}

// InsertAfter inserts a new element with a given value after
// a particular element.
func (l *List) InsertAfter(v int, prev *Element) *Element {
	e := NewElement(v) //creating new list element
	e.next = prev.next //setting prev el next el as new el next el
	if prev.next == nil {
		l.tail = e //if we insert last el, make it tail
	}
	prev.next = e //setting new el as prev el next el
	l.len++
	return e
}

//Converts linked list into slice
func ListToSlice(l *List) []int {
	var s []int
	s = append(s, int(l.head.Value))
	el := l.head.next
	for el != nil {
		s = append(s, int(el.Value))
		el = el.next
	}
	return s
}

//Converts slice into linked list using PushFront
func SliceToListFront(s []int) *List {
	l := &List{}
	for _, el := range s {
		l.PushFront(el)
	}
	return l
}

//Converts slice into linked list using PushBack
func SliceToListBack(s []int) *List {
	l := &List{}
	for _, el := range s {
		l.PushBack(el)
	}
	return l
}
