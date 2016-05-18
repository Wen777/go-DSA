package doublylinkedlist

import (
	"fmt"
)

// List represents a doubly linked list.
type List struct {
	first *element
	last  *element
	size  int
}

// element is a structure of link list
type element struct {
	value      interface{}
	prev, next *element
}

// New initialize a new empty list
func New() *List {
	return &List{}
}

// Add values into list
func (list *List) Add(values ...interface{}) {
	for _, item := range values {
		newElement := &element{value: item, prev: list.last}
		if list.size == 0 {
			list.first = newElement
			list.last = newElement
		} else {
			list.last.next = newElement
			list.last = newElement
		}
		list.size++
	}
}

// Append values into list
func (list *List) Append(values ...interface{}) {
	list.Add(values)
}

// Prepends insert node at the first index of list
func (list *List) Prepends(values ...interface{}) {
	for index := len(values) - 1; index >= 0; index-- {
		newElement := &element{value: values[index], next: list.first}
		if list.size == 0 {
			list.first = newElement
			list.last = newElement
		} else {
			list.first.prev = newElement
			list.first = newElement
		}
		list.size++
	}
}

// Get Returns the element at index.
// Second return parameter is true if index is within bounds of the array and array is not empty, otherwise false.
func (list *List) Get(index int) (interface{}, bool) {
	if !list.withinRange(index) {
		return nil, false
	}

	if list.size-index < index {
		elem := list.last
		element := list.last
		for i := 0; i < list.size-index; i, elem = i+1, elem.prev {
			fmt.Println("FIRST ", i)
		}
		for e := list.size - 1; e != index; e, element = e-1, element.prev {
			fmt.Println("SECOND ", e)
		}
		return elem.value, true
	}

	element := list.first
	for e := 0; e != index; e, element = e+1, element.next {
	}
	return element.value, true
}

// withinRange Check that the index is withing bounds of the list
func (list *List) withinRange(index int) bool {
	return index >= 0 && index < list.size && list.size != 0
}
