package singlylinkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	res := New()
	exp := &List{}
	assert.IsType(t, exp, res, "Expected type is same")
}

func TestPrepends(t *testing.T) {
	li := New()
	li.Append(3)
	li.Prepends(1999)
	assert.False(t, li.IsEmpty(), "Expected list isn't empty")
}

func TestNumberOf(t *testing.T) {
	li := New()
	li.Append(3)
	li.Append(10)
	li.Append(107)
	li.Append("vvv")

	assert.Equal(t, 4, li.NumberOf(), "Expected the number of node of list is 1")
}

func TestRemoveAll(t *testing.T) {
	li := New()
	li.Append(3)
	li.Append(10)
	li.Append(107)
	li.Append("vvv")

	li.RemoveAll()

	assert.Equal(t, 0, li.NumberOf(), "Expected the number of node of list is 0. There should be no any node in list.")
}

func TestGet(t *testing.T) {
	li := New()
	li.Append(3)
	li.Append(10)

	res, _ := li.Get(0)

	assert.Equal(t, 3, res, "Expected get same data")
}

func TestRemove(t *testing.T) {
	li := New()
	li.Append(3)
	li.Append(10)
	li.Append(107)

	li.Remove(1)
	res, _ := li.Get(1)

	assert.NotEqual(t, 10, res, "Expected 10 already removed from list")
}
