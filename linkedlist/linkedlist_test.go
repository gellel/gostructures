package linkedlist_test

import (
	"fmt"
	"testing"

	"github.com/gellel/gostructures/linkedlist"
)

func Test(t *testing.T) {

	l := linkedlist.New()

	fmt.Println(l.Add("a"))

	fmt.Println(l.Add("b"))

	l.Walk()

	fmt.Println(l.Search("b"))

	fmt.Println(l.Delete("b"))

	l.Walk()

	fmt.Println(l.Search("b"))
}
