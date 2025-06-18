package main

import (
	"fmt"
)

func PointersPassByValueToCopy() {
	// Pass by value to copy struct

	fmt.Printf("Copy without pointers is simple:\n\n")

	type person struct {
		Name string
	}
	rename := func(p person, name string) person {
		p.Name = name
		return p
	}
	alice := person{Name: "Alice"}
	fmt.Printf("Alice %v\n", alice)

	bob := rename(alice, "Bob")

	fmt.Printf("alice: %v, bob: %v \n", alice, bob)
}
