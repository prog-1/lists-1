package list

// Element is an element of a singly linked list.
type Element struct {
	// The value stored with this element.
	Value int
	next  *Element
}

// NewElement constructs and returns a new Element with a given value.
func NewElement(v int) *Element {
	return &Element{Value: v}
}

// Next returns the next list element or nil.
func (e *Element) Next() *Element {
	return e.next
}
