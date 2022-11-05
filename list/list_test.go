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
