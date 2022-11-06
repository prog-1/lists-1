package list

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPushback(t *testing.T) {
	l := List{}
	l.PushBack(0)
	//l.PushBack(1)
	l.PushBack(34)
	l.PushBack(-7)
	l.PushBack(0)
	if l.head.Value != 0 || l.head.next.Value != 34 || l.head.next.next.Value != -7 || l.tail.Value != 0 {
		t.Error("You are FAILURE! WORK MORE!")
		t.Error("Ahh, okey litle hint:")
		t.Errorf("head = %v, last = %v", l.head, l.tail)
	}
}

func TestFind(t *testing.T) {
	l := List{}
	l.PushFront(4)
	l.PushFront(0)
	l.PushFront(-4)
	l.PushFront(4)
	l.PushFront(667)
	a := l.Find(4)
	if a.Value != 4 {
		t.Error("Meow! Erros's here")
	}
	// how my code passed test from the first try
}

func TestInsertAfter(t *testing.T) {
	l := List{}
	l.PushFront(4)
	l.PushFront(0)
	l.PushFront(-4)
	l.PushFront(4)
	l.PushFront(667)
	l.InsertAfter(90, l.head)
	if l.head.next.Value != 90 {
		t.Error("Code does not pass the test")
		t.Errorf("head: %v, head.next: %v", l.head, l.head.next)
	}
}
func TestList(t *testing.T) {
	for _, tc := range []struct {
		name         string
		initF, wantF func() List
		comp         func(a, b List) bool
	}{
		{"Len0", list(), dontCare, wantLen(0)},
		{"Len5", list(1, 2, 3, 4, 5), dontCare, wantLen(5)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, want := tc.initF(), tc.wantF()
			if !tc.comp(got, want) {
				t.Errorf("got = %v, want = %v", format(got), format(want))
			}
		})
	}
}

func list(e ...int) func() List {
	return func() (l List) {
		for i := range e {
			l.PushBack(e[i])
		}
		return l
	}
}

func dontCare() List { return List{} }

func equal(a, b List) bool {
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

func wantLen(want uint) func(a, _ List) bool {
	return func(a, _ List) bool { return a.Len() == want }
}

func format(l List) string {
	var b bytes.Buffer
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Fprintf(&b, "->[%v]", e.Value)
	}
	b.WriteString("->|")
	return b.String()
}
