package main

import (
	"reflect"
	"testing"
)

func TestPushFront(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    []int
		want []int
	}{
		{"Push Front 1", []int{4, 6, 2, 7}, []int{7, 2, 6, 4}},
		{"Push Front 2", []int{6, 8, 1, 3}, []int{3, 1, 8, 6}},
	} {
		t.Run(tc.name, func(t *testing.T) {

			//Pushing elements front
			l := SliceToListFront(tc.s)
			if !reflect.DeepEqual(ListToSlice(l), tc.want) {
				t.Errorf("got = %v, want = %v", tc.s, tc.want)
			}
		})
	}
}

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
			if int(l.Find(T(tc.v)).next.Value) != tc.next {
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
			l.InsertAfter(T(tc.v), l.Find(T(tc.prev)))
			if !reflect.DeepEqual(ListToSlice(l), tc.want) {
				t.Errorf("got = %v, want = %v", tc.s, tc.want)
			}
		})
	}
}

func TestSortedInsert(t *testing.T) {
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
			l := SliceToListBack(tc.s)
			//Inserting value
			l.SortedInsert(T(tc.v))
			//Checking
			if !reflect.DeepEqual(ListToSlice(l), tc.want) {
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
		{"Insertion Sort 4", []int{1}, []int{1}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			//List from slice
			l := SliceToListBack(tc.s)
			//Checking
			if !reflect.DeepEqual(ListToSlice(InsertionSort(l)), tc.want) {
				Print(InsertionSort(l))
				t.Errorf("got = %v, want = %v", tc.s, tc.want)
			}
		})
	}
}
