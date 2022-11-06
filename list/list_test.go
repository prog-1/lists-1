package list

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func listtoslice(l *List) []int {
	var s []int
	for n := l.Front(); n != nil; n = n.Next() {
		s = append(s, n.Value)
	}
	return s
}
func slicetolistback(s []int) *List {
	var l *List
	for _, v := range s {
		l.PushBack(v)
	}
	return l
}
func slicetolistfront(s []int) *List {
	var l *List
	for _, v := range s {
		l.PushFront(v)
	}
	return l
}
func TestPushBack(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		n    int
		want []int
	}{

		{[]int{}, 1, []int{1}},
		{[]int{0, 120, 21}, 1, []int{0, 120, 21, 1}},
		{[]int{1, 2, 54, 101, 4, 1}, 25, []int{1, 2, 54, 101, 4, 1, 25}},
		{[]int{1, 2, 54, 101, 4, 1}, 3, []int{1, 2, 54, 101, 4, 1, 3}},
		{[]int{1, 2, -54, 101, 4, -1, 5}, 190, []int{1, 2, -54, 101, 4, -1, 5, 190}},
	} {
		l := slicetolistfront(tc.s)
		l.PushBack(tc.n)
		if got := listtoslice(l); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("find(%v,%v) = %v, want %v", tc.s, tc.n, got, tc.want)
		}
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
