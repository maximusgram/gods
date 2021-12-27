package singlylinkedlist

import (
	"github.com/maximusgram/gods/utils"
)

// List holds the elements, where each element points to the next element
type List struct {
	first *element
	last  *element
	size  int
}

type element struct {
	value interface{}
	next  *element
}

// New instantiates a new list and adds the passed values, if any, to the list
func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

// Add appends a value (one or more) at the end of the list (same as Append())
func (list *List) Add(values ...interface{}) {
	for _, value := range values {
		newElement := &element{value: value}
		if list.size == 0 {
			list.first = newElement
			list.last = newElement
		} else {
			list.last.next = newElement
			list.last = newElement
		}
		list.size++
	}
}

// Append appends a value (one or more) at the end of the list (same as Add())
func (list *List) Append(values ...interface{}) {
	list.Add(values...)
}

// Prepend prepends a value (or more)
func (list *List) Prepend(values ...interface{}) {
	for _, value := range values {
		newElement := &element{value: value, next: list.first}
		list.first = newElement
		if list.size == 0 {
			list.last = newElement
		}
		list.size++
	}
}

// Get returns the element at index
// Second return pararmeter is true if index is within bounds of the array and array is not empty, otherwise false
func (list *List) Get(index int) (interface{}, bool) {

	if !list.withinRange(index) {
		return nil, false
	}

	element := list.first
	for i := 0; i < index; i++ {
		element = element.next
	}

	return element.value, true
}

// Remove removes the element at the given index from the list
func (list *List) Remove(index int) {

	if !list.withinRange(index) {
		return
	}

	if list.size == 1 {
		list.Clear()
		return
	}

	var beforeElement *element
	curr_element := list.first

	for i := 0; i != index; i, curr_element = i+1, curr_element.next {
		beforeElement = curr_element
	}

	if curr_element == list.first {
		list.first = curr_element.next
	}

	if curr_element == list.last {
		list.last = beforeElement
	}

	if beforeElement != nil {
		beforeElement.next = curr_element.next
	}

	curr_element = nil
	list.size--
}

// Contains checks if values (one or more) are present in the set
// All values have to be present in the set for the method to return true
// Performance time complexity of n^2
// Returns true if no arguments are passed at all, i.e., set is always super-set of empty set
func (list *List) Contains(values ...interface{}) bool {
	if len(values) == 0 {
		return false
	}
	if list.size == 0 {
		return false
	}
	for _, value := range values {
		found := false
		for element := list.first; element != nil; element = element.next {
			if element.value == value {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Values returns all elements in the list
func (list *List) Values() []interface{} {
	values := make([]interface{}, list.size)
	for i, element := 0, list.first; i < list.size; i, element = i+1, element.next {
		values[i] = element.value
	}
	return values
}

// IndexOf returns index of the provided element
func (list *List) IndexOf(value interface{}) int {
	if list.size == 0 {
		return -1
	}
	for i, element := 0, list.first; i < list.size; i, element = i+1, element.next {
		if element.value == value {
			return i
		}
	}
	return -1
}

// Empty returns true if list does not contain any elements
func (list *List) Empty() bool {
	return list.size == 0
}

// Size returns number of elements within the list
func (list *List) Size() int {
	return list.size
}

// Clear removes all elements within the list
func (list *List) Clear() {
	list.size = 0
	list.first = nil
	list.last = nil
}

// Sort values
// use copied values to sort
// copy back the sorted values to list
func (list *List) Sort(comparator utils.Comparator) {
	if list.size < 2 {
		return
	}

	values := list.Values()
	utils.Sort(values, comparator)

	list.Clear()
	list.Add(values...)
}

func (list *List) withinRange(index int) bool {
	return index >= 0 && index < list.size
}
