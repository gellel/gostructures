package closed_test

import (
	"fmt"
	"testing"

	closed "github.com/gellel/gostructures/abstracts/hash-tables/closed-table"
)

func Test(t *testing.T) {

	table := &closed.Hash{}

	keys := []string{
		"GoLang",
		"Java",
		"JavaScript",
		"Lua",
		"Perl",
		"Python"}

	for i, key := range keys {
		fmt.Println(table.Add(key, i))
	}

	fmt.Println(table)

	for _, key := range keys {
		fmt.Println(table.Get(key))
	}

	for _, key := range keys {
		fmt.Println(table.Delete(key))
	}

	fmt.Println(table)
}
