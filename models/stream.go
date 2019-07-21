package models

import (
	"errors"
)

type Stream struct {
	Length int
	Limit  int
	start  *Element
	end    *Element
}

func (s *Stream) Append(newElement *Element) {
	if s.Length == 0 {
		s.start = newElement
		s.end = newElement
	} else {
		lastElement := s.end
		lastElement.next = newElement
		s.end = newElement
	}

	if s.Length == s.Limit {
		s.start = s.start.next
	} else {
		s.Length++
	}
}

func (s *Stream) Get(key string) (*Element, error) {
	if s.Length == 0 {
		return nil, errors.New("stream is empty")
	}

	currentElement := s.start

	for currentElement.Key != key {
		if currentElement.next == nil {
			return nil, errors.New("no such Element found")
		}
		currentElement = currentElement.next
	}

	return currentElement, nil
}

func (s *Stream) Remove(key string) {
	if s.Length == 0 {
		panic(errors.New("stream is empty"))
	}

	var previousElement *Element
	currentElement := s.start

	for currentElement.Key != key {
		if currentElement.next == nil {
			panic(errors.New("no such Element found"))
		}

		previousElement = currentElement
		currentElement = currentElement.next
	}
	previousElement.next = currentElement.next

	s.Length--
}

func (s *Stream) RemoveAll() {
	s.Length = 0
	s.start = nil
	s.end = nil
}
