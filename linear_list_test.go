package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetElem(t *testing.T) {
	l := LinearList{length: 6, data: [MAXSIZE]int{5, 6, 87, 2, 9, 50}}

	res, err := GetElem(l, 3)
	assert.Nil(t, err)
	assert.Equal(t, res, 87, "expected they should be equal")
}

func TestInsertList(t *testing.T) {
	num := 6
	arr := [MAXSIZE]int{5, 61, 78, 3, 44, 88}
	l := LinearList{length: num, data: arr}
	err := InsertList(&l, 2, 99)

	assert.Nil(t, err)
	assert.Equal(t, LinearList{length: 7, data: [MAXSIZE]int{5, 99, 61, 78, 3, 44, 88}}, l, "Expected that 99 insert to second position and length of list is 7")
}

func TestDeleteList(t *testing.T) {
	num := 6
	arr := [MAXSIZE]int{5, 61, 78, 3, 44, 88}
	l := LinearList{length: num, data: arr}
	var pop int
	// FIXME can't get the exact value of pop
	err := DeleteList(&l, 1, &pop)

	assert.Nil(t, err)
	assert.Equal(t, 5, pop, "Expected first index should be removed.")
	assert.Equal(t, LinearList{length: 5, data: [MAXSIZE]int{61, 78, 3, 44, 88}}, l, "Expected that remove first index and length of list is 5")
}
