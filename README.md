# LRU Cache Go Package

Implementation of LRU Cache in Go

[![Go Report Card](https://goreportcard.com/badge/github.com/miguelbemartin/lru-cache)](https://goreportcard.com/report/github.com/miguelbemartin/lru-cache)
[![Build Status](https://travis-ci.org/miguelbemartin/lru-cache.svg?branch=master)](https://travis-ci.org/miguelbemartin/lru-cache)
[![codecov](https://codecov.io/gh/miguelbemartin/lru-cache/branch/master/graph/badge.svg)](https://codecov.io/gh/miguelbemartin/lru-cache)
[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/miguelbemartin/lru-cache/master/LICENSE)

## Getting Started

### Prerequisites
- Go 1.10+

### Installation
Run into the terminal the next command

```
go get github.com/miguelbemartin/lru-cache
```

## Usage
```go
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


```

## Run the tests
```
go test ./... -v
```

## Contributing

## Authors
* **Miguel Ángel Martín** - [@miguelbemartin](https://twitter.com/miguelbemartin)

## License
This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
