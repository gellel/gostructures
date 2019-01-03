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

}
