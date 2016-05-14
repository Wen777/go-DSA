package list

import (
// "errors"
// "fmt"
)

// List represents a doubly linked list.
// The zero value for List is an empty list ready to use.
type List struct {
	head, last *Node
	len        int
}

// Node is a structure of link list
type Node struct {
	next *Node
	Val  interface{}
}

// Prepends insert node at the first index of list
func (list *List) Prepends(values ...interface{}) {
	for _, value := range values {
		newNode := &Node{next: list.head, Val: value}
		list.head = newNode
		if list.len == 0 {
			list.last = newNode
		}
		list.len++
	}
}

// Append a value (one or more) at the end of the list (same as Add())
func (list *List) Append(value interface{}) {
	newNode := &Node{Val: value}
	if list.len == 0 {
		list.head = newNode
		list.last = newNode
	} else {
		list.last.next = newNode
		list.last = newNode
	}
	list.len++
}

// IsEmpty returns true if list does not contain any elements.
func (list *List) IsEmpty() bool {
	return list.len == 0
}

// RemoveAll returns true if list does not contain any elements.
func (list *List) RemoveAll() bool {
	var p *Node
	p = list.head
	for p != nil && p.next != nil {
		p.next = p.next.next
	}

	// p = nil
	list.head = nil
	list.last = nil
	list.len = 0
	return true
}

// NumberOf show how many node in this list
func (list *List) NumberOf() int {
	var (
		num int
		p   *Node
	)
	p = list.head
	for p != nil {
		p = p.next
		num++
	}
	list.len = num
	return num
}

// withinRange Check that the index is withing bounds of the list
func (list *List) withinRange(index int) bool {
	return index >= 0 && index < list.len && list.len != 0
}

// Get Returns the element at index.
// Second return parameter is true if index is within bounds of the array and array is not empty, otherwise false.
func (list *List) Get(index int) (interface{}, bool) {

	if !list.withinRange(index) {
		return nil, false
	}

	element := list.head
	for e := 0; e != index; e, element = e+1, element.next {
	}

	return element.Val, true
}

// Remove , according gaven index to remove that node
func (list *List) Remove(index int) {
	if !list.withinRange(index) {
		return
	}
	if list.len == 1 {
		list.len = 0
		list.head = nil
		list.last = nil
	}
	var beforeNode *Node
	node := list.head
	for i := 0; i != index; i, node = i+1, node.next {
		beforeNode = node
	}

	if node == list.head {
		list.head = node.next
	}

	if node == list.last {
		list.last = beforeNode
	}

	if beforeNode != nil {
		beforeNode.next = node.next
	}

	node = nil
	list.len--

}

// New returns an initialized list.
func New() *List {
	l := &List{}
	l.len = 0
	return l
}
