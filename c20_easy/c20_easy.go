package c20_easy

import (
	"errors"
)

type Stack struct {
	val []rune
}

func (s *Stack) isEmpty() bool {
	return len(s.val) == 0
}

func (s *Stack) Pop() (rune, error) {
	if s.isEmpty() {
		return 0, errors.New("stack is empty")
	}
	if s.Len() == 1 {
		popVal := s.val[0]
		s.val = []rune{}
		return popVal, nil
	}
	popVal := s.val[len(s.val)-1]
	retainVal := s.val[0:len(s.val)-1]
	s.val = retainVal
	return popVal, nil
}

func (s *Stack) Push(val rune) {
	s.val = append(s.val, val)
}

func (s *Stack) Len() int {
	return len(s.val)
}

func isValid(s string) bool {
	var stack Stack
	mapping := map[rune]rune {
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if mapping[char] != 0 {
			topElement, err := stack.Pop()
			if err != nil {
				topElement = '#'
			}
			if mapping[char] != topElement {
				return false
			}
		} else {
			stack.Push(char)
		}
	}

	return stack.isEmpty()
}
