package doublylinkedlist

import (
	"fmt"
	"strings"

	"github.com/maximusgram/gods/utils"
)

// List holds the elements, where each element points to the next and previous element
type List struct {
    first *element
    last *element
    size int
}

type element struct {
    value interface{}
    prev *element
    next *element
}

// New instantiates a new list and adds the passed values, if any, to the list
func New(values ...interface{}) *List {
    return nil
}

// Add appends a value (one or more) at the end of the list (same as Append())
func (list *List) Add(values ...interface{}) {

}

// Append appends a value (one or more) at the end of the list (same as Add())
func (list *List) Append(values ...interface{}) {

}

// Prepend preprends a values (or more)
func (list *List) Prepend(values ...interface{}) {

}

// Remove removes the element at the given index from the list
func (list *List) Remove(index int) {

}

func (list *List) Contains(values ...interface{}) bool {
    return false
}

func (list *List) Values() []interface{} {
    return nil
}

func (list *List) IndexOf(value interface{}) int {
    return -1
}

func (list *List) Empty() bool {
    return false
}

func (list *List) Size() int {
    return -1
}

func (list *List) Clear() {

}

func (list *List) Sort(comparator utils.Comparator) {

}

func (list *List) Swap(i, j int) {

}

func (list *List) Insert(index int, values ...interface{}) {

}

func (list *List) Set(index int, value interface{}) {

}

func (list *List) Get(index int) (interface{}, bool) {
    return nil, false
}

func (list *List) String() string {
    str := "DoubleLinkedList\n"
    values := []string{}
    for element := list.first; element != nil; element = element.next {
        values = append(values, fmt.Sprintf("%v", element.value))
    }
    str += strings.Join(values, ", ")
    return str
}

func (list *List) withinRange(index int) bool {
    return index >= 0 && index < list.size
}







