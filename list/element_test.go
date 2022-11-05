package list

import "testing"

func TestElement(t *testing.T) {

	t.Run("NewElement", func(t *testing.T) {
		const want = 123
		e := NewElement(want)
		if e == nil {
			t.Fatal("got unexpected nil pointer")
		}
		if e.Value != want {
			t.Errorf(".Value = %v, want = %v", e.Value, want)
		}
		if e.Next() != nil {
			t.Errorf(".Next() = %v, want = nil", e.Next())
		}
	})

	t.Run("Next", func(t *testing.T) {
		e := NewElement(-1)
		const want = 123
		e.next = NewElement(want)
		if e.Next() == nil {
			t.Fatal("got unexpected nil pointer")
		}
		if got := e.Next().Value; got != want {
			t.Errorf(".Next().Value = %v, want = %v", got, want)
		}
	})
}
