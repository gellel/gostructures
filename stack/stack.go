package stack

import (
	"github.com/gellel/gostructures/stack/s"
)

// Stack of interfaces.
type Stack struct {
	s.Stack
	stack []interface{}
}

// New instantiates a new Queue pointer.
// Size argument controls size of Queue.
func New(size int) *Stack {
	return &Stack{stack: make([]interface{}, 0)}
}

// InStack checks that the provided int is in the Stack.stack range.
func (stack *Stack) InStack(i int) bool {
	return stack.IsEmpty() && i <= stack.SizeOf()
}

// IsEmpty checks that the stack contains entries.
func (stack *Stack) IsEmpty() bool {
	return stack.SizeOf() == 0
}

// Peek shows the item in  Stack.stack at the provided position. Does not pop the item.
func (stack *Stack) Peek(i int) interface{} {
	if stack.InStack(i) {
		return stack.stack[i]
	}
	return nil
}

// Pop takes the highest element off Stack.stack.
// Returns element and stack size if property is in range.
// Returns nil and -1 if empty.
func (stack *Stack) Pop() (interface{}, int) {
	if !stack.IsEmpty() {
		return stack.stack[:stack.SizeOf()-1], stack.SizeOf()
	}
	return nil, -1
}

// Push adds a new property onto the top of Stack.stack.
func (stack *Stack) Push(property interface{}) int {
	stack.stack = append(stack.stack, property)

	return stack.SizeOf()
}

// SizeOf shows the size of the current Stack.stack.
func (stack *Stack) SizeOf() int {
	return len(stack.stack)
}
