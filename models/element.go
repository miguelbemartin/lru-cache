package models

type Element struct {
	key   string
	value interface{}
	next  *Element
}
