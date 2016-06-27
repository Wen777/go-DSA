package arraylist

// import (
// 	"fmt"
// )

const (
	// GrowthFactor growth by 100%
	GrowthFactor = float32(2.0)
	// ShrinkFactor shrink when size is 25% of capacity (0 means never shrink)
	ShrinkFactor = float32(0.25)
)

// List represents a array linked list.
type List struct {
	elements []interface{}
	size     int
}

// New initialize a new empty list
func New() *List {
	return &List{}
}

// Add values into list
func (list *List) Add(elements ...interface{}) {
	list.growBy(len(elements))
	for _, element := range elements {
		list.elements[list.size] = element
		list.size++
	}
}

// reseize, based on input (expected capacity of array) to resize the length of array.
func (list *List) resize(capacity int) {
	newElements := make([]interface{}, capacity, capacity)
	copy(newElements, list.elements)
	list.elements = newElements
}

// growBy Extend list if the capacity isn't enough.
func (list *List) growBy(n int) {
	currentCapacity := cap(list.elements)
	if list.size+n >= currentCapacity {
		newCapacity := int(GrowthFactor * float32(currentCapacity+n))
		list.resize(newCapacity)
	}
}

// Shrink the array if necessary, i.e. when size is SHRINK_FACTOR percent of current capacity
func (list *List) shrink() {
	if ShrinkFactor == 0.0 {
		return
	}
	currentCapacity := cap(list.elements)
	if list.size <= int(float32(currentCapacity)*ShrinkFactor) {
		list.resize(list.size)
	}
}

// Get Returns the element at index.
// Second return parameter is true if index is within bounds of the array and array is not empty, otherwise false.
func (list *List) Get(index int) (interface{}, bool) {
	if !list.withinRange(index) {
		return nil, false
	}
	return list.elements[index], true
}

// Remove According index to remove specific element from list
func (list *List) Remove(index int) {
	if !list.withinRange(index) {
		return
	}
	list.elements[index] = nil
	copy(list.elements[index:], list.elements[index+1:list.size])
	list.size--
	list.shrink()
}

// Values return all elements in the list
func (list *List) Values() []interface{} {
	newElements := make([]interface{}, list.size, list.size)
	copy(newElements, list.elements[:list.size])
	return newElements
}

// Destroy remove all elements from list
func (list *List) Destroy() {
	list.elements = []interface{}{}
	list.size = 0
}

// IsEmpty if it is a empty list or not
func (list *List) IsEmpty() bool {
	return list.size == 0
}

// Size display size of list
func (list *List) Size() int {
	return list.size
}

// withinRange Check that the index is withing bounds of the list
func (list *List) withinRange(index int) bool {
	return index >= 0 && index < list.size && list.size != 0
}

// Swap values of two elements at the given indices.
func (list *List) Swap(i, j int) {
	if list.withinRange(i) && list.withinRange(j) {
		list.elements[i], list.elements[j] = list.elements[j], list.elements[i]
	}
}

// Contains Check if elements (one or more) are present in the set.
// All elements have to be present in the set for the method to return true.
// Performance time complexity of n^2.
// Returns true if no arguments are passed at all, i.e. set is always super-set of empty set.
func (list *List) Contains(elements ...interface{}) bool {
	for _, searchElement := range elements {
		found := false
		for _, element := range list.elements {
			if element == searchElement {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
