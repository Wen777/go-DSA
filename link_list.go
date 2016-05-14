package main

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
	list *List
}

// HeadInsert insert node at the begining of list
func (list *List) HeadInsert(value interface{}) {
	newNode := &Node{Val: value}
	if list.len == 0 {
		list.head = newNode
		list.last = newNode
	} else {
		newNode.next = list.head
		list.head = newNode
	}
	list.len++
}

// Append a value (one or more) at the end of the list (same as Add())
func (list *List) Append(value interface{}) {
	newNode := &Node{Val: value, list: list}
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

// New returns an initialized list.
func New() *List {
	l := &List{}
	l.len = 0
	return l
}
