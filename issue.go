package main

import (
	"fmt"
	"reflect"
)

type Ordering[T any] interface {
	// By commenting Compare(), the compiler does not fail terribly anymore.
	// It still fails gracefully, telling us about a wrong type for Reverse().

	Compare(a, b T) int

	Reverse() Ordering[T]
}

type Comp[T any] struct {
}

func (o Comp[T]) Compare(a, b T) int {
	return 0
}

func (o Comp[T]) Reverse() Ordering[T] {
	return o
}

func Lexicographical[T any]() Ordering[[]T] {
	return &Comp[[]T]{}
}

func main() {
	o := Lexicographical[string]()
	fmt.Println(reflect.TypeOf(o))
}
