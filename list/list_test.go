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
		fmt.Fprintf(&b, "->[%v]", e.Value)
	}
	b.WriteString("->|")
	return b.String()
}

func TestPushBack(t *testing.T) {
	for _, tc := range []struct {
		name         string
		initF, wantF func() List
		v            int
	}{
		{"case-1", list(), list(1), 1},
		{"case-2", list(1, 2, 3, 4, 5), list(1, 2, 3, 4, 5, 6), 6},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, want := tc.initF(), tc.wantF()
			got.PushBack(tc.v)
			if !equal(got, want) {
				t.Errorf("got = %v, want = %v", format(got), format(want))
			}
		})
	}
}

func TestFind(t *testing.T) {
	for _, tc := range []struct {
		name  string
		initF func() List
		v     int
		want  bool
	}{
		{"case-1", list(), 1, false},
		{"case-2", list(1, 2, 3), 3, true},
		{"case-3", list(1, 2, 3, 4, 5), 4, true},
		{"case-4", list(1, 2, 3), 5, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			l := tc.initF()
			got := l.Find(tc.v)
			if tc.want && got == nil {
				t.Errorf("got nil, want %v", tc.v)
			} else if tc.want && got.Value != tc.v {
				t.Errorf("got %v, want %v", got.Value, tc.v)
			} else if !tc.want && got != nil {
				t.Errorf("got %v, want nil", got.Value)
			}
		})
	}
}

func TestInsertAfter(t *testing.T) {
	for _, tc := range []struct {
		name         string
		initF, wantF func() List
		v, prev      int
	}{
		{"case-1", list(1, 2), list(1, 2, 3), 3, 2},
		{"case-2", list(1, 2, 4), list(1, 2, 3, 4), 3, 2},
	} {
		t.Run(tc.name, func(t *testing.T) {
			l, want := tc.initF(), tc.wantF()
			prev := l.Find(tc.prev)
			got := l.InsertAfter(tc.v, prev)
			if got.Value != tc.v {
				t.Errorf("got = %v, want = %v", got.Value, tc.v)
			} else if l != want {
				t.Errorf("got = %v, want = %v", format(l), format(want))
			}
		})
	}
}
