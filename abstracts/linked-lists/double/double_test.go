package double_test

import (
	"fmt"
	"testing"

	double "github.com/gellel/gostructures/abstracts/linked-lists/double"
)

func Test(t *testing.T) {

	list := double.New()

	list.Append("a").Append("b").Append("c")

	list.WalkReverse()

	fmt.Println("-")

	list.InsertBefore(list.Find("b"), "B")

	list.Walk()

	fmt.Println("-")

	list.InsertAfter(list.Find("b"), "C")

	list.Walk()

	fmt.Println("-")

	list.WalkReverse()

	fmt.Println("-")

	list.Remove("B")

	fmt.Println("-")

	list.Walk()
}
