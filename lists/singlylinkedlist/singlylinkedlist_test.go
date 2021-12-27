package singlylinkedlist

import "testing"

func TestListRemove(t *testing.T) {
    list := New()
    list.Add("a", "b", "c")
    list.Remove(2)
    if actualValue, ok := list.Get(1); actualValue == nil && !ok {
        t.Errorf("got %v, expected %v", actualValue, "b")
    }
}