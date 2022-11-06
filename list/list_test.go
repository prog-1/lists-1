package list

import (
	"bytes"
	"fmt"
	"testing"
)

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
		fmt.Fprintf(&b, "[%v]->", e.Value)
	}
	b.WriteString("|")
	return b.String()
}

func TestPushFront(t *testing.T) {
	for _, tc := range []struct {
		name  string
		list  func() List
		input int
		want  func() List
	}{
		{"1", list(1), 1, list(1, 1)},
		{"2", list(1), 2, list(2, 1)},
		{"3", list(2), 1, list(1, 2)},
		{"4", list(1, 2), 3, list(3, 1, 2)},
		{"5", list(1, 3), 2, list(2, 1, 3)},
		{"6", list(5, 3, 12, 9, 0, 5, 9, 2, 22, 1, 0, 90, 4, 6), 13, list(13, 5, 3, 12, 9, 0, 5, 9, 2, 22, 1, 0, 90, 4, 6)},
	} {
		t.Run(fmt.Sprint(tc.name), func(t *testing.T) {
			l := tc.list()
			l.PushFront(tc.input)
			if !equal(l, tc.want()) {
				t.Errorf("PushFront(%v %v) = %v, want %v", format(tc.list()), tc.input, format(l), format(tc.want()))
			}
		})
	}
}

func TestPushBack(t *testing.T) {
	for _, tc := range []struct {
		name  string
		list  func() List
		input int
		want  func() List
	}{
		{"1", list(1), 1, list(1, 1)},
		{"2", list(1), 2, list(1, 2)},
		{"3", list(2), 1, list(2, 1)},
		{"4", list(1, 2), 3, list(1, 2, 3)},
		{"5", list(1, 3), 2, list(1, 3, 2)},
		{"6", list(5, 3, 12, 9, 0, 5, 9, 2, 22, 1, 0, 90, 4, 6), 13, list(5, 3, 12, 9, 0, 5, 9, 2, 22, 1, 0, 90, 4, 6, 13)},
	} {
		t.Run(fmt.Sprint(tc.name), func(t *testing.T) {
			l := tc.list()
			l.PushBack(tc.input)
			if !equal(l, tc.want()) {
				t.Errorf("PushBack(%v %v) = %v, want %v", format(tc.list()), tc.input, format(l), format(tc.want()))
			}
		})
	}
}

func TestFind(t *testing.T) {
	for _, tc := range []struct {
		name  string
		list  func() List
		input int
	}{
		{"1", list(1), 1},
		{"2", list(1), 2},
		{"3", list(2), 2},
		{"4", list(2), 1},
		{"5", list(1, 1), 1},
		{"5", list(1, 2), 1},
		{"5", list(1, 2), 2},
		{"6", list(1, 2), 3},
		{"6", list(1, 3), 1},
		{"6", list(1, 3), 2},
		{"6", list(1, 2, 3), 2},
		{"6", list(1, 2, 3), 4},
		{"6", list(3, 1, 2, 9, 12, 0, 3), 12},
		{"7", list(5, 3, 12, 9, 0, 5, 9, 2, 22, 1, 0, 90, 4, 6), 9},
		{"7", list(5, 3, 12, 9, 0, 5, 9, 2, 22, 1, 0, 90, 4, 6), 13},
	} {
		t.Run(fmt.Sprint(tc.name), func(t *testing.T) {
			l := tc.list()
			got := l.Find(tc.input)
			if got != nil && got.Value != tc.input {
				t.Errorf("Find(%v %v) = %v, want %v", format(tc.list()), tc.input, got.Value, tc.input)
			} else if got == nil {
				for n := l.Front(); n != nil; n = n.Next() {
					if n.Value == tc.input {
						t.Errorf("Find(%v %v) = nil, want %v", format(tc.list()), tc.input, tc.input)
					}
				}
			}
		})
	}
}

func TestInsertAfter(t *testing.T) {
	for _, tc := range []struct {
		name         string
		list         func() List
		value, after int
		want         func() List
	}{
		{"1", list(1), 1, 1, list(1, 1)},
		{"2", list(1), 2, 1, list(1, 2)},
		{"3", list(2), 1, 2, list(2, 1)},
		{"4", list(1, 2), 3, 2, list(1, 2, 3)},
		{"5", list(1, 3), 2, 3, list(1, 3, 2)},
		{"5", list(1, 3), 2, 1, list(1, 2, 3)},
		{"6", list(5, 3, 12, 9, 0, 5, 9, 2, 22, 1, 0, 90, 4, 6), 13, 5, list(5, 13, 3, 12, 9, 0, 5, 9, 2, 22, 1, 0, 90, 4, 6)},
		{"6", list(5, 3, 12, 9, 0, 5, 9, 2, 22, 1, 0, 90, 4, 6), 256, 6, list(5, 3, 12, 9, 0, 5, 9, 2, 22, 1, 0, 90, 4, 6, 256)},
	} {
		t.Run(fmt.Sprint(tc.name), func(t *testing.T) {
			l := tc.list()
			l.InsertAfter(tc.value, l.Find(tc.after))
			if !equal(l, tc.want()) {
				t.Errorf("InsertAfter(%v %v %v) = %v, want %v", format(tc.list()), tc.value, tc.after, format(l), format(tc.want()))
			}
		})
	}
}

func TestSwapElements(t *testing.T) {
	for _, tc := range []struct {
		name            string
		list            func() List
		value, swapwith int
		want            func() List
	}{
		{"1", list(1, 2), 1, 2, list(2, 1)},
		{"1", list(1, 2), 2, 1, list(2, 1)},
		{"1", list(2, 1), 1, 2, list(1, 2)},
		{"5", list(1, 3), 1, 3, list(3, 1)},
		{"5", list(1, 3), 3, 1, list(3, 1)},
		{"5", list(3, 1), 3, 1, list(1, 3)},
		{"6", list(5, 3, 12, 9, 0, 5, 9, 2, 22, 1, 0, 90, 4, 6), 12, 4, list(5, 3, 4, 9, 0, 5, 9, 2, 22, 1, 0, 90, 12, 6)},
		{"6", list(5, 3, 12, 9, 0, 5, 9, 2, 22, 1, 0, 90, 4, 6), 3, 5, list(3, 5, 12, 9, 0, 5, 9, 2, 22, 1, 0, 90, 4, 6)},
	} {
		t.Run(fmt.Sprint(tc.name), func(t *testing.T) {
			l := tc.list()
			l.SwapElements(tc.value, tc.swapwith)
			if !equal(l, tc.want()) {
				t.Errorf("InsertAfter(%v %v %v) = %v, want %v", format(tc.list()), tc.value, tc.swapwith, format(l), format(tc.want()))
			}
		})
	}
}
