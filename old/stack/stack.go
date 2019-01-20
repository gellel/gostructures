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
func New() *Stack {
	return &Stack{stack: make([]interface{}, 0)}
}

// InStack checks that the provided int is in the Stack.stack range.
func (stack *Stack) InStack(i int) bool {
	return stack.IsPopulated() && i < stack.SizeOf()
}

// IsEmpty checks that Stack.stack does not contain entries.
func (stack *Stack) IsEmpty() bool {
	return stack.SizeOf() == 0
}

// IsPopulated checks that Stack.stack contains entries.
func (stack *Stack) IsPopulated() bool {
	return stack.SizeOf() > 0
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
		element := stack.next()
		return element, stack.drop()
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

func (stack *Stack) drop() int {
	stack.stack = stack.stack[:stack.SizeOf()-1]

	return stack.SizeOf()
}

func (stack *Stack) next() interface{} {
	return stack.stack[stack.SizeOf()-1]
}
