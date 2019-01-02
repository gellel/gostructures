package stack_test

import (
	"fmt"
	"testing"

	"github.com/gellel/gostructures/stack"
)

func Test(t *testing.T) {

	s := stack.New()

	fmt.Println(s)

	fmt.Println("push", s.Push(1))

	fmt.Println("size", s.SizeOf())

	fmt.Println("empty", s.IsEmpty())

	fmt.Println("peek", s.Peek(0))

	fmt.Println(s.Pop())
}
