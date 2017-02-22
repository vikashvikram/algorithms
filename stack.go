package compute

import (
	"errors"
	"fmt"
)

type Stack struct {
	data []interface{}
	length int
}

func (s *Stack) Pop() (interface{}, error) {
	if s.length == 0 {
		return 0, errors.New("Error: Stack is empty")
	}
	val := s.data[s.length-1]
	s.data = s.data[:s.length-1]
	s.length -= 1
	return val, nil
}

func (s *Stack) Push(val interface{})  {
	s.data = append(s.data, val)
	s.length += 1
}

func (s *Stack) Top() interface{}  {
	return s.data[len(s.data)-1]
}

func (s *Stack) String() string {
	return fmt.Sprintf("%v", s.data)
}
