package main

import (
	"fmt"

	lrucache "github.com/miguelbemartin/lru-cache"
)

func main() {

	// Init the LRU Cache client
	myCache, err := lrucache.NewLRUCache(2)
	if err != nil {
		panic("Error")
	}

	// Add a key and value to cache
	err = myCache.Set("my-key", "my value")
	if err != nil {
		panic("Error adding a new element to cache")
	}

	// Add a key and value to cache
	err = myCache.Set("my-key-2", "my value 2")
	if err != nil {
		panic("Error adding a new element to cache")
	}

	// Add a key and value to cache
	err = myCache.Set("my-key-3", "my value 3")
	if err != nil {
		panic("Error adding a new element to cache")
	}

	// Fetch value by key
	value, err := myCache.Get("my-key")
	if err != nil {
		panic("Error fetching new value from cache")
	}

	// Clean cache
	err = myCache.Prune()
	if err != nil {
		panic("Error cleaning the cache")
	}

	// Print the value
	fmt.Println("My value is: ", value)
}
