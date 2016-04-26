package main

import (
	"errors"
	"fmt"
)

// MAXSIZE : declare capacity for linear list
const MAXSIZE = 20

// LinearList is a structure of linear list
// FIXME this struct expect the type of data must be int.
// In the future, implement support differet type of data
type LinearList struct {
	length int          // length of list
	data   [MAXSIZE]int //storage of data
}

// GetElem according integer to find where is the value
func GetElem(list LinearList, i int) (int, error) {
	if list.length == 0 || i < 0 || i > list.length {
		return 0, errors.New("unexpected input")
	}
	return list.data[i-1], nil
}

// InsertList insert new value. i is index.
func InsertList(list *LinearList, i int, value int) error {
	if list.length == MAXSIZE {
		return errors.New("list is full")
	}
	if i < 1 || i > list.length+1 {
		return errors.New("invalidate input")
	}

	if i <= list.length {
		for k := list.length - 1; k >= i-1; k-- {
			list.data[k+1] = list.data[k]
		}
	}
	list.data[i-1] = value
	list.length++
	return nil
}

// DeleteList created for removing value
func DeleteList(list *LinearList, i int, popValue *int) error {
	if list.length == 0 {
		return errors.New("list is empty")
	}
	if i < 1 || i > list.length {
		return errors.New("index is invalidate")
	}
	*popValue = list.data[i-1]
	if i < list.length {
		for k := i; k < list.length; k++ {
			list.data[k-1] = list.data[k]
		}
	}
	list.data[list.length-1] = 0
	list.length--
	return nil
}

func main() {
	// initialize
	l := LinearList{length: 6, data: [MAXSIZE]int{5, 6, 87, 2, 9, 50}}

	// GetElem
	fmt.Println("excute linear list initialized")
	res, err := GetElem(l, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("GetElem of l in 2,", res)

	// Linear list insertion
	fmt.Println("Before insert, linear list is", l)
	err = InsertList(&l, 2, 999)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("After insert, linear list is", l)
}
