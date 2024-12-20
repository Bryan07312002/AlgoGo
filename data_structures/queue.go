package datastructures

import "errors"

type Queue[T any] []T

// Appends a item to the final of the Queue
func (q *Queue[T]) push(item T) {
	*q = append(*q, item)
}

// Returns the next item of the Queue without removeing it
func (q *Queue[T]) peek() (T, error) {
	if len(*q) == 0 {
		var item T
		return item, errors.New("peek called when Queue is empty")
	}

	return (*q)[0], nil
}

// Returns next item and removes it from the Queue
func (q *Queue[T]) pop() (T, error) {
	var item T
	if len(*q) == 0 {
		return item, errors.New("pop called when Queue is empty")
	}

	item = (*q)[0]
	*q = (*q)[1:]

	return item, nil
}
