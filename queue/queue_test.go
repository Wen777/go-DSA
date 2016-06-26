package queue

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestPush(t *testing.T) {
	q := NewQueue()

	q.Push(5)
	q.Push("GGG")
	assert.Equal(t, 5, q.head.data, "Expect the data of haed of Queue is 5")
	assert.Equal(t, "GGG", q.tail.data, "Expect the data of haed of Queue is 'GGG'")
	assert.Equal(t, 2, q.Len(), "Expect the length of Queue is 2")

}

func TestPushNil(t *testing.T) {
	q := NewQueue()

	for i := 0; i < 15; i++ {
		q.Push(nil)
	}
	assert.Equal(t, 15, q.Len(), "Expect the length of Queue is 15")
	for i := 15; i > 0; i-- {
		assert.Equal(t, i, q.Len(), "Expect the length of Queue decreasing")
		assert.Equal(t, nil, q.Pop(), "Expect the length of Queue decreasing")
	}

}

func TestPop(t *testing.T) {
	q := NewQueue()

	assert.Equal(t, nil, q.Pop(), "Expect the value of oldest node of queue is nil")

	q.Push(5)
	q.Push("GGG")

	data1 := q.Pop()
	data2 := q.Pop()

	assert.Equal(t, 5, data1, "Expect the data of haed of Queue is 5")
	assert.Equal(t, "GGG", data2, "Expect the data of haed of Queue is 'GGG'")
	assert.Equal(t, 0, q.Len(), "Expect the length of Queue is 0")

}

func TestPeek(t *testing.T) {
	q := NewQueue()

	assert.Equal(t, nil, q.Peek(), "Expect the value of oldest node of queue is nil")

	q.Push(5)
	q.Push("GGG")

	data1 := q.Peek()
	data2 := q.Peek()

	assert.Equal(t, 5, data1, "Expect the data of haed of Queue is 5")
	assert.Equal(t, data1, data2, "Expect the value is same")

}

func TestLen(t *testing.T) {
	q := NewQueue()

	assert.Equal(t, 0, q.Len(), "Expect length of Queue is 0")
}

//--Concurrent tests

func TestConcurrent(t *testing.T) {
	q := NewQueue()
	numberGoRoutines := 50
	numberOfPushes := 10000
	var wg sync.WaitGroup
	for i := 0; i < numberGoRoutines; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < numberOfPushes; j++ {
				q.Push(j)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	assert.Equal(t, numberGoRoutines*numberOfPushes, q.Len(), "Expected length of queue is product of numberGoRoutines and numberOfPushes ")
}
