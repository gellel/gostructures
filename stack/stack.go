package stack

import (
	"github.com/gellel/gostructures/stack/s"
)

type Stack struct {
	s.Stack
	stack []interface{}
}

func (stack *Stack) InStack(i int) bool {
	return stack.IsEmpty() && i <= stack.SizeOf()
}

func (stack *Stack) IsEmpty() bool {
	return stack.SizeOf() == 0
}

func (stack *Stack) Peek(i int) interface{} {
	if stack.InStack(i) {
		return stack.stack[i]
	}
	return nil
}

func (stack *Stack) Push(property interface{}) int {
	stack.stack = append(stack.stack, property)

	return stack.SizeOf()
}

func (stack *Stack) SizeOf() int {
	return len(stack.stack)
}
