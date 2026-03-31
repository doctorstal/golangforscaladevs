package main

import "fmt"

type Foo interface {
	Bar(string)
}

type AnonymousFoo struct{}

func (AnonymousFoo) Bar(string) {
	fmt.Println("I do not need access arguments to run!")
}

func NewFoo() Foo {
	return AnonymousFoo{}
}

// NOTE: underscore _ is stil used as Blank Identifier
