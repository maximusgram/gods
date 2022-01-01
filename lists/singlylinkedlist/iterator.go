package singlylinkedlist

type Iterator struct {
	list    *List
	index   int
	element *element
}

func (list *List) Iterator() Iterator {
	return Iterator{list: list, index: -1, element: nil}
}

func (it *Iterator) Next() bool {
	if it.index < it.list.size {
		it.index++
	}
	if !it.list.withinRange(it.index) {
		return false
	}
	if it.index == 0 {
		it.element = it.list.first
	} else {
        it.element = it.element.next
	}
    return true
}

func (it *Iterator) Value() interface{} {
	return it.element.value
}

func (it *Iterator) Index() int {
	return it.index
}

func (it *Iterator) Begin() {
    it.index = -1
    it.element = nil
}

func (it *Iterator) First() bool {
    it.Begin()
    return it.Next()
}
