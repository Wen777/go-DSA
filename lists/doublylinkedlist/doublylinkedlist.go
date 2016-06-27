package doublylinkedlist

import (
	"fmt"
)

// List represents a doubly linked list.
type List struct {
	head, last *element
	len        int
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
		if list.len == 0 {
			list.head = newElement
			list.last = newElement
		} else {
			list.last.next = newElement
			list.last = newElement
		}
		list.len++
	}
}

// Append values into list
func (list *List) Append(values ...interface{}) {
	list.Add(values)
}

// Prepends insert node at the head index of list
func (list *List) Prepends(values ...interface{}) {
	for index := len(values) - 1; index >= 0; index-- {
		newElement := &element{value: values[index], next: list.head}
		if list.len == 0 {
			list.head = newElement
			list.last = newElement
		} else {
			list.head.prev = newElement
			list.head = newElement
		}
		list.len++
	}
}

// Get Returns the element at index.
// Second return parameter is true if index is within bounds of the array and array is not empty, otherwise false.
func (list *List) Get(index int) (interface{}, bool) {
	if !list.withinRange(index) {
		return nil, false
	}

	if list.len-index < index {
		elem := list.last
		element := list.last
		for i := 0; i < list.len-index; i, elem = i+1, elem.prev {
			fmt.Println("HEAD ", i)
		}
		for e := list.len - 1; e != index; e, element = e-1, element.prev {
			fmt.Println("SECOND ", e)
		}
		return elem.value, true
	}

	element := list.head
	for e := 0; e != index; e, element = e+1, element.next {
	}
	return element.value, true
}

// withinRange Check that the index is withing bounds of the list
func (list *List) withinRange(index int) bool {
	return index >= 0 && index < list.len && list.len != 0
}
