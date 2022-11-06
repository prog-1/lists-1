package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/prog-1/list"
)

func List(e ...int) func() list.List {
	return func() (l list.List) {
		for i := range e {
			l.PushBack(e[i])
		}
		return l
	}
}

func equal(a, b list.List) bool {
	if a.Len() != b.Len() {
		return false
	}
	for a, b := a.Front(), b.Front(); a != nil; a, b = a.Next(), b.Next() {
		if a.Value != b.Value {
			return false
		}
	}
	return true
}

func format(l list.List) string {
	var b bytes.Buffer
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Fprintf(&b, "[%v]->", e.Value)
	}
	b.WriteString("|")
	return b.String()
}

func TestInsertionSort(t *testing.T) {
	for _, tc := range []struct {
		name        string
		input, want func() list.List
	}{
		{"1", List(1), List(1)},
		{"2", List(2), List(2)},
		{"3", List(1, 2), List(1, 2)},
		{"4", List(2, 1), List(1, 2)},
		{"5", List(1, 3), List(1, 3)},
		{"6", List(3, 1), List(1, 3)},
		{"7", List(1, 2, 3), List(1, 2, 3)},
		{"8", List(1, 3, 2), List(1, 2, 3)},
		{"9", List(2, 1, 3), List(1, 2, 3)},
		{"10", List(2, 3, 1), List(1, 2, 3)},
		{"11", List(3, 2, 1), List(1, 2, 3)},
		{"12", List(3, 1, 2), List(1, 2, 3)},
		{"12", List(1, 2, 1), List(1, 1, 2)},
		{"13", List(5, 3, 12, 9, 0, 5, 9, 2, 22, 1, 0, 90, 4, 6), List(0, 0, 1, 2, 3, 4, 5, 5, 6, 9, 9, 12, 22, 90)},
	} {
		t.Run(fmt.Sprint(tc.name), func(t *testing.T) {
			l := tc.input()
			less := func(a, b int) bool { return a > b }
			InsertionSort(&l, less)
			if !equal(l, tc.want()) {
				t.Errorf("PushFront(%v %v) = %v, want %v", format(tc.input()), format(tc.input()), format(l), format(tc.want()))
			}
		})
	}
}

func TestInsertSorted(t *testing.T) {
	//t.Error("must be implemented")
}
