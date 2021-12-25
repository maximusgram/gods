package containers

type EnumerableWithIndex interface {
    // Each calls the given function once for each element, passing that element's index and value
    Each(func(index int, value interface{}))

    // Map invokes the given function once for each element and returns 
    // a container containing the values returned by the given function
    // Map(func(index int, value interface{}) interface{}) Container

    // Select returns a new container containing all elements for which the given function returns a 
    // true value
    // Select(func(index int, value interface{}) bool) Container

    // Any passes each element of the container to the given function and
    // return true if the function ever returns true for any element
    Any(func(index int, value interface{}) bool) bool

    // Find passes each element of the the container of the given function and returns
    // the first (index, value) for which the function is true or -1, nil otherwise
    // if no element matches the criteria.
    Find(func(index int, value interface{}) bool) (int, interface{})
}

// EnumberableWithKey provides functions for ordered containers whose values/ elements are key/value pairs.
type EnumberableWithKey interface {
    // Each calls the given function for each element, passing that element's key and value
    Each(func(key interface{}, value interface{}))

    // Map invokes the given function once for each element and returns 
    // a container containing the values returned by the given function as key/value pairs
    // Map(func(key interface{}, value interface{}) (interface{}, interface{}) Container

    // Select returns a new container containing all elements for which the given function returns a 
    // true value
    // Select(func(key interface{}, values interface{}) bool) Container

    // Any passes each element of the container to the given function and
    // returns true if the function ever returns true for any element
    Any(func(key interface{}, value interface{}) bool) bool

    // Find passes each element of the container to the given function and returns 
    // the first (key, value) for which the function is true or nil, nil otherwise if 
    // no element matches the criteria
    Find(func(key interface{}, value interface{}) bool) (interface{}, interface{})
}