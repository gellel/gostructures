package single_test

import (
	"fmt"
	"testing"

	single "github.com/gellel/gostructures/abstracts/linked-lists/single"
)

func Test(t *testing.T) {

	list := single.New()

	fmt.Println(list.Find("1"))

	fmt.Println(list)

	list.Append("a").Append("b").Append("c")

	fmt.Println(list.Find("b"))

	fmt.Println(list.Size())

	list.InsertAfter(list.Find("b"), "b.1")

	fmt.Println(list.InsertBefore(list.Find("b.1"), "b.0"))

	list.Walk()
}
