package stack

import (
	"sync"
)

type node struct {
	data interface{}
	next *node
}

// Stack a type of classical data structure
type Stack struct {
	head *node
	lock *sync.Mutex
	len  int
}

// NewStack return a instance of Stack
func NewStack() *Stack {
	s := &Stack{}
	s.lock = &sync.Mutex{}
	return s
}

// Len return the length of stack
// concurrent safty
func (s *Stack) Len() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.len
}

// Push insert a data into Stack
// cocurrent safty
func (s *Stack) Push(data interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	n := &node{data: data}
	if s.head != nil {
		n.next = s.head
	}
	s.head = n

	s.len++
}

// Pop take out a data from Stack
// cocurrent safty
func (s *Stack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.head == nil {
		return nil
	}

	n := s.head
	s.head = n.next
	s.len--

	return n.data
}

// Peek show a data of newest node from Stack
// cocurrent safty
func (s *Stack) Peek() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.head == nil {
		return nil
	}

	return s.head.data
}
