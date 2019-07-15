package lrucache

import "errors"

type client struct {
	Size   int
	Memory map[string]interface{}
}

func NewLRUCache(size int) (*client, error) {
	c := &client{
		Size:   size,
		Memory: make(map[string]interface{}),
	}

	return c, nil
}

func (c *client) Set(key string, value string) error {

	if (len(c.Memory) >= c.Size) {

	}

	c.Memory[key] = value

	return nil
}

func (c *client) Get(key string) (interface{}, error) {
	value, ok := c.Memory[key]
	if !ok {
		return nil, errors.New("element not found")
	}

	return value, nil
}

func (c *client) Prune() error {
	for k := range c.Memory {
		delete(c.Memory, k)
	}

	return nil
}
