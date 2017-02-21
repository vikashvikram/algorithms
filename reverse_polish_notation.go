package algorithms

import (
	"os"
	"fmt"
	"errors"
	"strconv"
)

var (
	input_arr = os.Args[1:]
	rpn_stack = &stack{[]int{}, 0}
)

type stack struct {
	data []int
	length int
}

func (s *stack) Pop() (int, error) {
	if s.length == 0 {
		return 0, errors.New("Error: Stack is empty")
	}
	val := s.data[s.length-1]
	s.data = s.data[:s.length-1]
	s.length -= 1
	return val, nil
}

func (s *stack) Push(val int)  {
	s.data = append(s.data, val)
	s.length += 1
}

func ReversePolishNotation() (int, error)  {
	for i := 0; i < len(input_arr); i++ {
		char := input_arr[i]
		doCal(char)
	}
	return rpn_stack.Pop()
}

func doCal(char string)  {
	switch char {
	case "+":
		add()
	case "-":
		subtract()
	case "x":
		multiply()
	case "/":
		divide()
	default:
		if val, err := strconv.Atoi(char); err == nil {
			rpn_stack.Push(val)
		} else {
			fmt.Println(err)
		}
	}
}

func getArguments() (int, int)  {
	val2, _ := rpn_stack.Pop()
	val1, _ := rpn_stack.Pop()
	return val1, val2
}

func add()  {
	val1, val2 := getArguments()
	rpn_stack.Push(val1 + val2)
}

func subtract()  {
	val1, val2 := getArguments()
	rpn_stack.Push(val1 - val2)
}

func multiply()  {
	val1, val2 := getArguments()
	rpn_stack.Push(val1 * val2)
}

func divide()  {
	val1, val2 := getArguments()
	rpn_stack.Push(val1 / val2)
}
