package models

import (
	"fmt"
	"testing"
)

func TestStream_Append(t *testing.T) {
	var tests = []struct {
		elements     []Element
		sizeExpected int
		limit        int
	}{
		{
			elements: []Element{
				{
					Key:   "key1",
					Value: "value 1",
				},
				{
					Key:   "key2",
					Value: "value 2",
				},
			},
			sizeExpected: 2,
			limit:        10,
		},
		{
			elements: []Element{
				{
					Key:   "key1",
					Value: "value 1",
				},
				{
					Key:   "key2",
					Value: "value 2",
				},
				{
					Key:   "key3",
					Value: "value 3",
				},
			},
			sizeExpected: 3,
			limit:        10,
		},
		{
			elements: []Element{
				{
					Key:   "key1",
					Value: "value 1",
				},
				{
					Key:   "key2",
					Value: "value 2",
				},
			},
			sizeExpected: 1,
			limit:        1,
		},
	}

	for i, test := range tests {
		stream := Stream{
			Limit:  test.limit,
			Length: 0,
		}

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			for _, element := range test.elements {
				stream.Append(&element)
			}
			if stream.Length != test.sizeExpected {
				t.Errorf("Test failed: wrong size: %d\ngot:  %d", test.sizeExpected, stream.Length)
			}
		})
	}
}

func TestStream_Get(t *testing.T) {
	var tests = []struct {
		elements      []Element
		limit         int
		key           string
		errorExpected string
	}{
		{
			elements: []Element{
				{
					Key:   "key1",
					Value: "value 1",
				},
				{
					Key:   "key2",
					Value: "value 2",
				},
			},
			limit:         1,
			key:           "key1",
			errorExpected: "no such element found",
		},
		{
			elements: []Element{
				{
					Key:   "key3",
					Value: "value 3",
				},
				{
					Key:   "key4",
					Value: "value 4",
				},
				{
					Key:   "key5",
					Value: "value 5",
				},
			},
			limit:         10,
			key:           "key4",
			errorExpected: "",
		},
	}

	for i, test := range tests {

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {

			stream := Stream{
				Limit:  test.limit,
				Length: 0,
			}

			for _, element := range test.elements {
				newElement := Element{
					Key:   element.Key,
					Value: element.Value,
					next:  nil,
				}
				stream.Append(&newElement)
			}

			_, err := stream.Get(test.key)

			if err != nil && err.Error() != test.errorExpected {
				t.Errorf("Test failed: wrong error: %q\ngot:  %q", err.Error(), test.errorExpected)
			}
		})
	}
}
