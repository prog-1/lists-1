package main

import (
	"github.com/prog-1/list"

	"testing"
)

func TestInsertionSort(t *testing.T) {
	l := list.List{}
	var less LessFunc
	l.PushBack(1)
	l.PushBack(3)
	l.PushBack(4)
	l.PushFront(2)
	InsertionSort(&l, less) //why &&&&&&
	if l.Front().Value != 1 {
		t.Error("i am tired of your mistakes, go to sleeep")
	}
}

func TestInsertSorted(t *testing.T) {
	l := list.List{}
	l.PushBack(1) // lets just pretend insertion sort works
	l.PushBack(2)
	l.PushBack(3)
	SortedInsert(&l, 9, func(a, b int) bool { return a < b }) // tests cant run with this &? idk
	if l.Front().Value != 1 {
		t.Error("i am tired of your mistakes, go to sleeep")
	}
}
