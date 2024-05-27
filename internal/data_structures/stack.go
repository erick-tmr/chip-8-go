package data_structures

import "errors"

var ErrStackOverflow = errors.New("data_structures-stack: overflow, too many elements")

type Stack32[T any] [32]T

var pointer uint8

func (s *Stack32[T]) Push(val T) error {
	if pointer == 32 {
		return ErrStackOverflow
	}

	s[pointer] = val
	pointer++
	return nil
}

func (s *Stack32[T]) Pop() (T, bool) {
	if pointer == 0 {
		var zero T
		return zero, false
	}

	top := s[pointer-1]
	var zero T
	s[pointer-1] = zero
	pointer--

	return top, true
}
