package main

import "fmt"

// Type system: structs, receiver functions, interfaces, avoid interface pointers

// simple struct
type privateStruct struct {
	privateAttribute string
}
type PublicStruct struct {
	PublicAttribute string
}

var (
	privateVar int
	PublicVar  int
)

type person struct {
	Name string
	Age  int
}

var p = &person{}

// Receiver function, i.e. "method"
func (p *person) rename(newName string) {
	p.Name = newName
}

// Could be also value receiver
func (p person) copy() person {
	return p
}

// But usually pointer receiver 1. allow modification, 2. avoid copying

// Can check for nil - providing default value
func (p *person) NameOrDefault(defaultName string) string {
	if p != nil {
		return p.Name
	} else {
		return defaultName
	}
}

// String implements fmt.Stringer
func (p *person) String() string {
	return fmt.Sprintf("Person name: %s, age: %d", p.Name, p.Age)
}

var stringer fmt.Stringer = p

type TimeConsumer interface {
	AddOneYear()
}

// AddOneYear implements TimeConsumer.
func (p *person) AddOneYear() {
	panic("unimplemented")
}

var timeConsumer TimeConsumer = p
