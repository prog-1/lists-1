package main

import (
	"reflect"
	"testing"

	"github.com/prog-1/list"
)

func TestInsertSorted(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    []int
		v    int
		want []int
	}{
		{"Sorted Insert 1", []int{1, 2, 3, 5}, 4, []int{1, 2, 3, 4, 5}},
		{"Sorted Insert 2", []int{10, 20, 30, 40}, 25, []int{10, 20, 25, 30, 40}},
		{"Sorted Insert 3", []int{10, 20, 30, 40}, 5, []int{5, 10, 20, 30, 40}},
		{"Sorted Insert 3", []int{}, 5, []int{5}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			//List from slice
			l := list.SliceToListBack(tc.s)
			//Inserting value
			SortedInsert(l, tc.v)
			//Checking
			if !reflect.DeepEqual(list.ListToSlice(l), tc.want) {
				t.Errorf("got = %v, want = %v", tc.s, tc.want)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    []int
		want []int
	}{
		{"Insertion Sort 1", []int{2, 1, 3, 5, 4}, []int{1, 2, 3, 4, 5}},
		{"Insertion Sort 2", []int{2, 2, 1, 5, 4}, []int{1, 2, 2, 4, 5}},
		{"Insertion Sort 3", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		//{"Insertion Sort 4", []int{1}, []int{1}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			//List from slice
			l := list.SliceToListBack(tc.s)

			if !reflect.DeepEqual(list.ListToSlice(InsertionSort(l)), tc.want) {
				list.Print(InsertionSort(l))
				t.Errorf("got = %v, want = %v", tc.s, tc.want)
			}
		})
	}
}
