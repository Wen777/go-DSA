package doublylinkedlist

import (
	// "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	res := New()
	exp := &List{}
	assert.IsType(t, exp, res, "Expected type is same")
}

func TestPrepend(t *testing.T) {
	li := New()
	li.Append(3)
	li.Prepends(1999)
	assert.NotZero(t, li.size, "Expected list isn't empty")
}

func TestGet(t *testing.T) {
	li := New()
	li.Append(3)
	li.Append(10)

	relay, _ := li.Get(0)
	assert.Equal(t, []interface{}{3}, relay, "Expected get same data")
}
