package queue

import (
	"sync"
)

// node
type node struct {
	data interface{}
	next *node
}

// Queue one of classic data structure
type Queue struct {
	head, tail *node
	len        int
	lock       *sync.Mutex
	// waiters  waiters
	// items    items
	// disposed bool
}

// NewQueue function to create Queue instance
func NewQueue() *Queue {
	q := &Queue{}
	q.lock = &sync.Mutex{}
	return q
}

// Len retun length of queue
func (q *Queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.len
}

// Push insert new node into queue
func (q *Queue) Push(data interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	n := &node{data: data}
	if q.tail == nil {
		q.head = n
	} else {
		q.tail.next = n
	}
	q.tail = n
	q.len++
}

// Pop return the value of newest node of the queue
func (q *Queue) Pop() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.head == nil {
		return nil
	}

	n := q.head
	q.head = n.next
	q.len--
	return n.data
}

// Peek return the value of oldest node of the queue
func (q *Queue) Peek() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.head == nil {
		return nil
	}

	return q.head.data
}
