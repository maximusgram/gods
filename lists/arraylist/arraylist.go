package arraylist

import (
	"fmt"
	"strings"

	"github.com/maximusgram/gods/utils"
)

type List struct {
	elements []interface{}
	size     int
}

const (
	growthFactor = float32(2.0)  // growth by 100%
	shrinkFactor = float32(0.25) // shrink when size is 25% of capacity (0 menas never shrink)
)

// New instantiates a new list and adds the passed values, if any, to the list
func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

// Add appends a values at the end of the list
func (list *List) Add(values ...interface{}) {
	list.growBy(len(values))

	// fmt.Fprintf(writer, "Size of list %v", cap(list.elements))
	// return
	for _, value := range values {
		list.elements[list.size] = value
		list.size++
	}
}

// Get returns the element at index.
// Second return parameter is true if index is within bounds of the array and array is not empty, otherwise false
func (list *List) Get(index int) (interface{}, bool) {
	if !list.withinRange(index) {
		return nil, false
	}

	return list.elements[index], true
}

func (list *List) Remove(index int) {
	if !list.withinRange(index) {
		return
	}

	list.elements[index] = nil // cleanup reference
	copy(list.elements[index:], list.elements[index+1:list.size])
	list.size--

	list.shrink()
}

// Contains checks if elements (one or more) are present in the set.
// All elements have to be present in the set for the method to return true
// Perfrormance time complexity is n^2
// Returns true if no arguments are passed at all, i.e., set is always super-set of empty set.
func (list *List) Contains(values ...interface{}) bool {

	for _, searchValue := range values {
		found := false
		for _, element := range list.elements {
			if element == searchValue {
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

// Values returns all elements in the list
func (list *List) Values() []interface{} {
	newElements := make([]interface{}, list.size)
	copy(newElements, list.elements)
	return newElements
}

// IndexOf returns index of provided element
func (list *List) IndexOf(value interface{}) int {
	for index, element := range list.elements {
		if element == value {
			return index
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

// Clear removes all the elements from the list
func (list *List) Clear() {
	list.size = 0
	list.elements = []interface{}{}
}

// Sort sorts values (in-place)
func (list *List) Sort(comparator utils.Comparator) {
	if list.size < 2 {
		return
	}
	utils.Sort(list.elements[:list.size], comparator)
}

// Swap swaps the two values at the specified positions
func (list *List) Swap(i, j int) {
	if list.withinRange(i) && list.withinRange(j) {
		list.elements[i], list.elements[j] = list.elements[j], list.elements[i]
	}
}

// Insert inserts values at specified index position shifting the values at that position (if any) and any subsequent elements to the right
// Does not do anything if position is negative or bigger than list's size
// Note: position equal to list;s size is valid, i.e., append
func (list *List) Insert(index int, values ...interface{}) {

	if !list.withinRange(index) {
		// Append
		if index == list.size {
			list.Add(values...)
		}
		return
	}

	l := len(values)
	list.growBy(l)
	list.size += l
	copy(list.elements[index+l:], list.elements[index:list.size-l]) // shifting to the right by l elements
	copy(list.elements[index:], values)
}

// Set the value at specified index
// Does not do anything if position is negatie or bigger than list's size
// Note: position equal to list's size is valid, i.e., append
func (list *List) Set(index int, value interface{}) {
	if !list.withinRange(index) {
		// Append
		if index == list.size {
			list.Add(value)
		}
		return
	}

	list.elements[index] = value
}

// String returns a string representation of container
func (list *List) String() string {
	str := "Arraylist\n"
	values := []string{}
	for _, value := range list.elements {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

// Check that the index is within bounds of the list
func (list *List) withinRange(index int) bool {
	return index >= 0 && index < list.size
}

// Expand the array if necessary, i.e. capacity will be reached if we add n elements
func (list *List) growBy(n int) {
	// When capacity is reached, grow by a factor of growthFactor and add number of elements
	currentCapacity := cap(list.elements)
	if list.size+n >= currentCapacity {
		newCapacity := int(growthFactor * float32(currentCapacity+n)) 
		list.resize(newCapacity)
	}
}

func (list *List) resize(cap int) {
	newElements := make([]interface{}, cap)
	copy(newElements, list.elements)
	list.elements = newElements
}

func (list *List) shrink() {
	if shrinkFactor == 0.0 {
		return
	}

	// Shrink when size is at shrinkFactor * capacity
	currentCapacity := cap(list.elements)
	if list.size <= int(float32(currentCapacity)*shrinkFactor) {
		list.resize(list.size)
	}
}
