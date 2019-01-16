package single_test

import (
	"fmt"
	"testing"

	single "github.com/gellel/gostructures/abstracts/linked-lists/single"
)

func Test(t *testing.T) {

	list := single.New()

	list.Append("a").Append("b").Append("c")

	fmt.Println(list.Find("b"))

	fmt.Println(list.Size())

	list.InsertAfter(list.Find("b"), "b.1")

	list.InsertBefore(list.Find("b.1"), "b.0")

	list.Prepend(1)

	list.Prepend(0)

	list.Append(list)

	list.Walk()

	fmt.Println(list.Find("1"), list.Tail)

	list.Remove("b.1")

	list.Walk()
}
