package main

import (
	"fmt"
	"reflect"
	"testing"
)

type T int

type Element struct {
	Value T
	next  *Element
}

func NewElement(v T) *Element {
	return &Element{Value: v}
}

func (e *Element) Next() *Element {
	return e.next
}

type List struct {
	head, tail *Element
	len        uint
}

func (l *List) Front() *Element {
	return l.head
}

func (l *List) Back() *Element {
	return l.tail
}

func Print(l *List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
		if e.Next() != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}

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

func (l *List) PushBack(v T) {
	e := NewElement(v)
	if l.len == 0 {
		l.head, l.tail = e, e
	} else {
		l.tail.next = e //changing link of current last element to new element
		l.tail = e      //changing last element to new one
	}
	l.len++
}

func (l *List) Find(v T) *Element {
	el := l.head
	// for el != nil {
	// 	if el.Value == v {
	// 		return el
	// 	} else {
	// 		el = el.next
	// 	}
	// }
	//return el

	//return a == b, yeah? readability
	for el != nil && el.Value != v {
		el = el.next
	}
	return el
}

func (l *List) InsertAfter(v T, prev *Element) *Element {
	//we need to find element..
	e := NewElement(v) //creating new list element
	e.next = prev.next //setting prev el next el as new el next el
	prev.next = e      //setting new el as prev el next el
	l.len++
	return e
}

func (l *List) SortedInsert(v T) *Element {
	el := l.head
	for el.next != nil && el.next.Value < v {
		el = el.next
	}
	return l.InsertAfter(v, el)
}

func InsertionSort(l *List) *List {
	//TODO
	return nil
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

func SliceToListFront(s []int) *List {
	l := &List{}
	for _, el := range s {
		l.PushFront(T(el))
	}
	return l
}

func SliceToListBack(s []int) *List {
	l := &List{}
	for _, el := range s {
		l.PushBack(T(el))
	}
	return l
}

// func main() {
// 	l := &List{}
// 	l.PushFront(25)
// 	l.PushBack(3)
// 	l.PushFront(60)
// 	Print(l)
// 	fmt.Println(ListToSlice(l))
// }

func TestPushFront(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    []int
		want []int
	}{
		{"Push Front 1", []int{4, 6, 2, 7}, []int{7, 2, 6, 4}},
		{"Push Front 2", []int{6, 8, 1, 3}, []int{3, 1, 8, 6}},
	} {
		t.Run(tc.name, func(t *testing.T) {

			//Pushing elements front
			l := SliceToListFront(tc.s)
			if !reflect.DeepEqual(ListToSlice(l), tc.want) {
				t.Errorf("got = %v, want = %v", tc.s, tc.want)
			}
		})
	}
}

func TestPushBack(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    []int
		want []int
	}{
		{"Push Back 1", []int{4, 6, 2, 7}, []int{4, 6, 2, 7}},
		{"Push Back 2", []int{6, 8, 1, 3}, []int{6, 8, 1, 3}},
	} {
		t.Run(tc.name, func(t *testing.T) {

			//Pushing elements back
			l := SliceToListBack(tc.s)
			if !reflect.DeepEqual(ListToSlice(l), tc.want) {
				t.Errorf("got = %v, want = %v", tc.s, tc.want)
			}
		})
	}
}

func TestFind(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    []int
		v    int
		next int
	}{
		{"Find 1", []int{4, 6, 2, 7}, 6, 2},
		{"Find 2", []int{6, 8, 1, 3}, 1, 3},
	} {
		t.Run(tc.name, func(t *testing.T) {

			//List from slice
			l := SliceToListBack(tc.s)

			//comparing value of the next el of found el
			if int(l.Find(T(tc.v)).next.Value) != tc.next {
				t.Errorf("got = %v, want = %v", tc.s, tc.next)
			}
		})
	}
}

func TestInsertAfter(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    []int
		v    int
		prev int
		want []int
	}{
		{"Insert After 1", []int{4, 6, 2, 7}, 100, 6, []int{4, 6, 100, 2, 7}},
		{"Insert After 2", []int{6, 8, 1, 3}, 12, 1, []int{6, 8, 1, 12, 3}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			//List from slice
			l := SliceToListBack(tc.s)
			l.InsertAfter(T(tc.v), l.Find(T(tc.prev)))
			if !reflect.DeepEqual(ListToSlice(l), tc.want) {
				t.Errorf("got = %v, want = %v", tc.s, tc.want)
			}
		})
	}
}

func TestSortedInsert(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    []int
		v    int
		want []int
	}{
		{"Sorted Insert 1", []int{1, 2, 3, 5}, 4, []int{1, 2, 3, 4, 5}},
		{"Sorted Insert 2", []int{10, 20, 30, 40}, 25, []int{10, 20, 25, 30, 40}},
		//{"Sorted Insert 3", []int{10, 20, 30, 40}, 50, []int{10, 20, 30, 40, 50}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			//List from slice
			l := SliceToListBack(tc.s)
			//Inserting value
			l.SortedInsert(T(tc.v))
			//Checking
			if !reflect.DeepEqual(ListToSlice(l), tc.want) {
				t.Errorf("got = %v, want = %v", tc.s, tc.want)
			}
		})
	}
}

func main() {

	testing.Main(
		/* matchString */ func(a, b string) (bool, error) { return a == b, nil },
		/* tests */ []testing.InternalTest{
			{Name: "Test Push Front", F: TestPushFront},
			{Name: "Test Push Back", F: TestPushBack},
			{Name: "Test Find", F: TestFind},
			{Name: "Test Insert After", F: TestInsertAfter},
			{Name: "Test Sorted Insert", F: TestSortedInsert},
		},
		/* benchmarks */ nil /* examples */, nil)
}
