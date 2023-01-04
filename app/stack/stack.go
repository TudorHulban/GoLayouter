package stack

import (
	"errors"
	"strings"
)

type Stack []any

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(item any) {
	*s = append(*s, item)
}

func (s *Stack) Peek() any {
	if len(*s) == 0 {
		return nil
	}

	return (*s)[len(*s)-1]
}

func (s *Stack) Pop() any {
	if len(*s) == 0 {
		return errors.New("stack is empty")
	}

	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return res
}

func (s Stack) String() string {
	var res []string

	for _, item := range s {
		res = append(res, item.(string))
	}

	return strings.Join(res, "/")
}
