package main_test

import (
	"reflect"
	"testing"

	main "github.com/gschauer/go-issue-compiler-generics-slice"
)

func TestIssue(t *testing.T) {
	o := main.Lexicographical[string]()
	n := reflect.TypeOf(o).String()
	if n != "*main.Comp[[]string]" {
		t.Fatalf("\nwant \"*main.Comp[[]string]\"\ngot  \"%v\"", n)
	}
}
