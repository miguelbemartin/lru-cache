package lrucache

import (
	"errors"
	"github.com/miguelbemartin/lru-cache/models"
)

type client struct {
	size   int
	stream *models.Stream
}

func NewLRUCache(size int) (*client, error) {

	// Check if the size is higher than 0
	if size <= 0 {
		return nil, errors.New("size should be higher than 0")
	}

	stream := models.Stream{
		Length: 0,
		Limit:  size,
	}

	c := &client{
		size:   size,
		stream: &stream,
	}

	return c, nil
}

func (c *client) Set(key string, value interface{}) error {
	newElement := models.Element{
		Key:   key,
		Value: value,
	}

	c.stream.Append(&newElement)

	return nil
}

func (c *client) Get(key string) (interface{}, error) {
	value, err := c.stream.Get(key)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (c *client) Prune() error {
	c.stream.RemoveAll()
	return nil
}
