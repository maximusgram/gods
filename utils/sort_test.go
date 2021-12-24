package utils

import "testing"

func TestSortStrings(t *testing.T) {
    strings := []interface{}{}

    strings = append(strings, "d")
    strings = append(strings, "a")
    strings = append(strings, "b")
    strings = append(strings, "c")
    strings = append(strings, "e")

    Sort(strings, StringComparator)

    for i := 1; i < len(strings); i++ {
        if strings[i - 1].(string) > strings[i].(string) {
            t.Errorf("Not sorted!")
        }
    }
}

func TestSortStructs(t *testing.T) {

    type User struct {
        id int
        name string
    }

    byId := func(a, b interface{}) int {
        c1 := a.(User)
        c2 := b.(User)

        switch {
        case c1.id > c2.id:
            return 1
        case c1.id < c2.id:
            return -1
        default:
            return 0
        }
    }

    users := []interface{}{
        User{4, "a"},
        User{1, "b"}, 
        User{3, "c"}, 
        User{0, "d"}, 
        User{6, "e"}, 
    }

    Sort(users, byId)

    for i := 1; i < len(users); i++ {
        if users[i - 1].(User).id > users[i].(User).id {
            t.Error("Not sorted")
        }
    }
}