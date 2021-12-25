package containers

// IteratorWithIndex is stateful iterator for ordered containers whose values can be fetched by an index
type IteratorWithIndex interface {
	Next() bool
	Value() interface{}
	Index() int
	Begin()
	First() bool
}

/// IteratorWithKey is a stateful iterator for ordered containers whose elements are key values pairs
type IteratorWithKey interface {
	Next() bool
	Value() interface{}
	Key() interface{}
	Begin()
	First() bool
}

// Same as IteratorWithIndex, with additional functionalities
type ReverseIteratorWithIndex interface {
	Prev() bool
	End()
	Last() bool

	IteratorWithIndex
}

// Same as IteratorWithKey, with additional functionalities
type ReverIteratorWithKey interface {
	Prev() bool
	End()
	Last() bool

	IteratorWithKey
}

/*

---------------------- Standard Iterator Functions -----------------------

Next:
    Next movies the iterator to the next element and returns true if there was n next element in the container
    If Next() returns true, the next element's [index / key] and value can be retrieved by [Index() / Key()] and Value().
    If Next() was called for the first time, then it will point the iterator to the first element if it exists
    Modified the state of the iterator

Value:
    Value returns the current element's value.
    Does not modify the state of the iterator

Index:
    Index returns the current element's index.
    Does not modify the state of the iterator

Key:
    Key returns the current element's key.
    Does not modify the state of the iterator

Begin:
    Begin resets the iterator to its initial state (one-before-first)
    Call Next() to fetch the first element if any.

First:
    First moves the iterator to the first element and returns true if there was a first element in the container.
    If First() returns true, then first element's [index / key] and value can be retrieved by [Index() / Key()] and Value().
    Modifies the state of the iterator


---------------------- Reversed Iterator Functions -----------------------

Prev:
    Prev moves the iterator to the previous element and returns true if there was a previous element in the container.
    If Prev() returns true, then previous element's [index/ key] and value can be retrieved by [Index() / Key()] and Value().
    Modifies the state of the iterator

End:
    End moves the iterator past the last element (one-past-the-end)
    Call Prev() to fetch the last element if any

Last():
    Last moves the iterator to the last element and returns true if there was a last element in the container.
    If Last() returns true, the last element's [index / key] and value can be retrieved by Key() and Value()

*/
