package containers

import (
	"github.com/maximusgram/gods/utils"
	"testing"
)

// for testing purposes
type ContainerTest struct {
	values []interface{}
}

func (c *ContainerTest) Values() []interface{} {
	return c.values
}

func (c *ContainerTest) Size() int {
	return len(c.values)
}

func (c *ContainerTest) Empty() bool {
	return c.Size() == 0
}

func (c *ContainerTest) Clear() {
	c.values = []interface{}{}
}


func TestGetSortedValuesInts(t *testing.T) {
	container := &ContainerTest{}
	container.values = []interface{}{5, 1, 2, 3, 4}

	values := GetSortedValues(container, utils.IntComparator)
	for i := 1; i < container.Size(); i++ {
		if values[i-1].(int) > values[i].(int) {
			t.Error("Not sorted")
		}
	}

    container.Clear()
    if container.Size() > 0 {
        t.Errorf("Clear not successful! Size : %v", container.Size())
    }
}