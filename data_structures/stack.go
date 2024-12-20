package datastructures

import "errors"

type Stack[T any] []T

func (s *Stack[T]) push(value T) {
	*s = append(*s, value)
}

func (s *Stack[T]) pop() (T, error) {
	if len(*s) == 0 {
		var zeroVal T
		return zeroVal, errors.New("")
	}

	top := (*s)[len(*s)-1]

	*s = (*s)[0:len(*s)-1]

	return top, nil
}
