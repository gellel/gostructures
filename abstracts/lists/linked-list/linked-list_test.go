package single_test

import (
	"fmt"
	"testing"

	single "github.com/gellel/gostructures/abstracts/lists/linked-list"
)

func Test(t *testing.T) {

	list := single.New()

	fmt.Println(list)

	list.Append("a").Append("b").Append("c")

	fmt.Println(list.Find("b"))

	fmt.Println(list.Delete("b"))

	fmt.Println(list.SizeOf())

	fmt.Println(list.Contains("b"))

	fmt.Println(list.Prepend(1))

	fmt.Println(list.Head)
}
