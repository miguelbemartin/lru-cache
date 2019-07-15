package models

import "errors"

type Stream struct {
	length int
	start  *Element
	end    *Element
}

func (s *Stream) Append(newElement *Element) {
	if s.length == 0 {
		s.start = newElement
		s.end = newElement
	} else {
		lastElement := s.end
		lastElement.next = newElement
		s.end = newElement
	}
	s.length++
}

func (s *Stream) Remove(key string) {
	if s.length == 0 {
		panic(errors.New("stream is empty"))
	}

	var previousElement *Element
	currentElement := s.start

	for currentElement.key != key {
		if currentElement.next == nil {
			panic(errors.New("no such Element found"))
		}

		previousElement = currentElement
		currentElement = currentElement.next
	}
	previousElement.next = currentElement.next

	s.length--
}
