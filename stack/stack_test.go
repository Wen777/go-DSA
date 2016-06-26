package stack

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack()
	assert.Equal(t, 0, s.Len(), "Expect the length of Queue is 0")
}

func TestPush(t *testing.T) {
	s := NewStack()
	s.Push(99)

	assert.Equal(t, 1, s.Len(), "Expect the length of Queue is 1")
}

func TestPop(t *testing.T) {
	s := NewStack()
	assert.Equal(t, nil, s.Pop(), "Expect the value of stack is nil")
	assert.Equal(t, 0, s.Len(), "Expect the length of Queue is 0")

	s.Push(99)
	assert.Equal(t, 1, s.Len(), "Expect the length of Queue is 1")
	data := s.Pop()
	assert.Equal(t, 99, data, "Expect the value of stack is 99")
	assert.Equal(t, 0, s.Len(), "Expect the length of Queue is 0")

}

func TestPeek(t *testing.T) {
	s := NewStack()
	assert.Equal(t, nil, s.Peek(), "Expect the value of stack is nil")
	assert.Equal(t, 0, s.Len(), "Expect the length of Queue is 0")

	s.Push(99)
	assert.Equal(t, 1, s.Len(), "Expect the length of Queue is 1")
	data := s.Peek()
	assert.Equal(t, 99, data, "Expect the value of stack is 99")
	assert.Equal(t, 1, s.Len(), "Expect the length of Queue is 1")
}

//--Concurrent tests

func TestConcurrent(t *testing.T) {
	s := NewStack()
	numberGoRoutines := 50
	numberOfPushes := 10000
	var wg sync.WaitGroup
	for i := 0; i < numberGoRoutines; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < numberOfPushes; j++ {
				s.Push(j)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	assert.Equal(t, numberGoRoutines*numberOfPushes, s.Len(), "Expected length of queue is product of numberGoRoutines and numberOfPushes ")
}
