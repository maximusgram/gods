package singlylinkedlist

func (list *List) Each(f func(index int, value interface{})) {
    for element, index := list.first, 0; index < list.size && element != nil; element, index = element.next, index + 1 {
        f(index, element.value)
    }
}

func (list *List) Map(f func(index int, value interface{}) interface{}) *List{
    newList := New()
    for element, index := list.first, 0; index < list.size && element != nil; element, index = element.next, index + 1 {
        newList.Add(f(index, element.value))
    }
    return newList
}

func (list *List) Select(f func(index int, value interface{}) bool) *List {
    newList := New()
    for element, index := list.first, 0; index < list.size && element != nil; element, index = element.next, index + 1 {
        if f(index, element.value) {
            newList.Add(element.value)
        }
    }
    return newList
}

func (list *List) Any(f func(index int, value interface{}) bool) bool {
    for element, index := list.first, 0; index < list.size && element != nil; element, index = element.next, index + 1 {
        if f(index, element.value) {
            return true
        }
    }
    return false
}

func (list *List) All(f func(index int, value interface{}) bool) bool {
    for element, index := list.first, 0; index < list.size && element != nil; element, index = element.next, index + 1 {
        if !f(index, element.value) {
            return false
        }
    }
    return true
}

func (list *List) Find(f func(index int, value interface{}) bool) (index int, value interface{}) {
    for element, index := list.first, 0; index < list.size && element != nil; element, index = element.next, index + 1 {
        if f(index, element.value) {
            return index, element.value
        }
    }
    return -1, nil
}