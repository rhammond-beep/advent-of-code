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

/*
Pop and element from the stack, or if no elements on stack, return
a populated error message
*/
func (s *Stack[SE]) Pop() (*SE, error) {
	if len(s.data) == 0 {
		return nil, errors.New("Stack Empty")
	}

	n := len(s.data) - 1
	element := s.data[n]
	s.data = s.data[:n]
	return &element, nil
}

func TestStack() {
	var s Stack[int]
	s.Push(1)
	s.Push(2)
	s.Push(3)

	for {
		val, error := s.Pop()
		if error != nil {
			break
		}
		fmt.Println(*val)
	}

	var strStack Stack[string]
	strStack.Push("world")
	strStack.Push("hello")

	for {
		val, error := strStack.Pop()
		if error != nil {
			break
		}
		fmt.Println(*val)
	}
}
