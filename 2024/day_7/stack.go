package day7

import (
	"errors"
	"fmt"
)

type StackElement interface {
	int | int8 | int16 | rune | int64 | string | bool
}

type StackImpl[SE StackElement] interface {
	Pop() SE
	Push(se SE)
}

type Stack[SE StackElement] struct {
	data []SE
}

/*
Push new element onto the end of the stack
*/
func (s *Stack[SE]) Push(se SE) {
	s.data = append(s.data, se)
}

func (s *Stack[SE]) Pop() (*SE, error) {
	n := len(s.data) - 1
	if n == -1 {
		return nil, errors.New("Stack Empty")
	}
	element := s.data[n]
	s.data = s.data[:n]
	return &element, nil
}

func TestStack() {
	var s Stack[int]

	// s.Push(1)
	// s.Push(2)
	// s.Push(3)

	val, error := s.Pop()

	if error != nil {
		fmt.Println(error)
	} else {

		fmt.Println(*val)
	}
}
