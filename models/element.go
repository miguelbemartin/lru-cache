package models

type Element struct {
	Key   string
	Value interface{}
	next  *Element
}
