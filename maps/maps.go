package maps

import "github.com/maximusgram/gods/containers"

// Map interface that all maps implement
type Map interface {
    Put(key interface{}, value interface{})
    Get(key interface{}, value interface{}, found bool)
    Remove(key interface{})
    Keys() []interface{}

    containers.Container
}

