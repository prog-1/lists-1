package list

import (
	"reflect"
	"testing"
)

// func TestList(t *testing.T) {
// 	t.Error("must be implemented")
// }

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
			if int(l.Find(tc.v).next.Value) != tc.next {
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
			l.InsertAfter(tc.v, l.Find(tc.prev))
			if !reflect.DeepEqual(ListToSlice(l), tc.want) {
				t.Errorf("got = %v, want = %v", tc.s, tc.want)
			}
		})
	}
}
