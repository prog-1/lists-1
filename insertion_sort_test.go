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
	InsertionSort(&l, less)
	if l.Front().Value != 1 {
		t.Error("i am tired if your mistakes, go to sleeep")
	}
}

func TestInsertSorted(t *testing.T) {
	t.Error("must be implemented")
}
