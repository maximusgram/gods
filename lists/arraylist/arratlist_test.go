package arraylist

import (
	"bytes"
	"fmt"
	"testing"
)

func TestLIstNew(t *testing.T) {
	list1 := New()

	if observed, expected := list1.Size(), 0; observed != expected {
		t.Errorf("expected %v got %v", expected, observed)
	}

	buffer := bytes.Buffer{}
	list1.Add(1, 2, 3, 4, 5)
    fmt.Print(buffer.String())

	if observed, expected := list1.Size(), 5; observed != expected {
		t.Errorf("expected %v got %v | %v", expected, observed, buffer.String())
	}
}
