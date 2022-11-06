package main

import (
	"reflect"
	"testing"

	"github.com/prog-1/list"
)

func TestInsertionSort(t *testing.T) {
	for _, tc := range []struct {
		name string
		init []int
		want []int
	}{
		{"1", []int{4, 2, 1, 3}, []int{1, 2, 3, 4}},
		{"empty", []int{}, []int{}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Initializing list
			var l list.List
			for i := len(tc.init) - 1; i >= 0; i-- {
				l.PushFront(tc.init[i])
			}

			InsertionSort(&l, func(a, b int) bool { return a < b })

			got := list.ListToSlice(&l)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got = %v, want = %v", l, tc.want)
			}
		})
	}
}

func TestSortedInsert(t *testing.T) {
	for _, tc := range []struct {
		name  string
		init  []int
		input int
		want  []int
	}{
		{"empty", []int{}, 0, []int{0}},
		{"last", []int{0, 1}, 2, []int{0, 1, 2}},
		{"first", []int{1, 2}, 0, []int{0, 1, 2}},
		{"middle", []int{0, 1, 3}, 2, []int{0, 1, 2, 3}},
		{"dublicate", []int{0, 1, 2, 3}, 2, []int{0, 1, 2, 2, 3}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Initializing list
			var l list.List
			for i := len(tc.init) - 1; i >= 0; i-- {
				l.PushFront(tc.init[i])
			}

			// Calling tested function:
			SortedInsert(&l, tc.input, func(a, b int) bool { return a < b })

			if !reflect.DeepEqual(list.ListToSlice(&l), tc.want) {
				t.Errorf("got = %v, want = %v", l, tc.want)
			}
		})
	}
}
