package compute

import (
	"os"
	"errors"
	"strconv"
	"fmt"
)

var (
	input_arr = os.Args[1:]
	rpn_stack = &Stack{}
)

func ReversePolishNotation() (int, error)  {
	for i := 0; i < len(input_arr); i++ {
		char := input_arr[i]
		doCal(char)
	}
	if val, err := rpn_stack.Pop(); err != nil {
		return 0, err
	} else {
		return val.(int), nil
	}
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
	return val1.(int), val2.(int)
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
