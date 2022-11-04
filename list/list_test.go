package list

import "testing"

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
