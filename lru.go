package lrucache

import (
	"github.com/miguelbemartin/lru-cache/models"
)

type client struct {
	size   int
	stream *models.Stream
}

func NewLRUCache(size int) (*client, error) {
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
