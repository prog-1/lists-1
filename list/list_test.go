package list

import (
	"reflect"
	"testing"
)

func TestPushBack(t *testing.T) {
	for _, tc := range []struct {
		name  string
		init  []int
		input []int
		want  []int
	}{
		{"empty", []int{}, []int{0}, []int{0}},
		{"1", []int{1, 2}, []int{3}, []int{1, 2, 3}},
		{"2", []int{1, 2}, []int{3, 4}, []int{1, 2, 3, 4}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// List initialization:
			var l List
			for i := len(tc.init) - 1; i >= 0; i-- {
				l.PushFront(tc.init[i])
			}

			// Calling tested function:
			for _, v := range tc.input {
				l.PushBack(v)
			}

			if got := ListToSlice(&l); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestFind(t *testing.T) {
	for _, tc := range []struct {
		name  string
		init  []int
		input int
		want  int
	}{
		{"empty", []int{}, 0, -1},
		{"found", []int{0, 1}, 1, 1},
		{"not found", []int{0, 1, 2}, 3, -1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// List initialization:
			var l List
			for i := len(tc.init) - 1; i >= 0; i-- {
				l.PushFront(tc.init[i])
			}

			// Calling tested function:
			output := l.Find(tc.input)
			var got int
			switch output {
			case nil:
				got = -1
			default:
				got = output.Value
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestInsertAfter(t *testing.T) {
	type Input struct{ v, prevVal int }
	for _, tc := range []struct {
		name  string
		init  []int
		input Input
		want  []int
	}{
		{"last", []int{0, 1, 2}, Input{3, 2}, []int{0, 1, 2, 3}},
		{"middle", []int{0, 1}, Input{2, 0}, []int{0, 2, 1}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// List initialization:
			var l List
			for i := len(tc.init) - 1; i >= 0; i-- {
				l.PushFront(tc.init[i])
			}

			// Calling tested function:
			prev := l.Find(tc.input.prevVal)
			output := l.InsertAfter(tc.input.v, prev)

			// Checking list order:
			if got := ListToSlice(&l); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("incorrect list order: got = %v, want = %v", got, tc.want)
			}
			// Checking output correctness:
			if output.Value != tc.input.v {
				t.Errorf("incorrect output: got = %v, want = %v", output.Value, tc.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	for _, tc := range []struct {
		name  string
		init  []int
		input int
		want  []int
	}{
		{"first", []int{0, 1, 2}, 0, []int{1, 2}},
		{"last", []int{0, 1, 2}, 2, []int{0, 1}},
		{"middle", []int{0, 1, 2}, 1, []int{0, 2}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// List initialization:
			var l List
			for i := len(tc.init) - 1; i >= 0; i-- {
				l.PushFront(tc.init[i])
			}

			l.Remove(l.Find(tc.input))

			if got := ListToSlice(&l); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
