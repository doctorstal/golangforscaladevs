package main

import (
	"fmt"
)

// Mutability and pointers
func PointersExample() {
	// Basics
	var x int = 10
	var ptr *int = &x

	fmt.Print("Basics\n\n")
	fmt.Println(x)    // output: 10
	fmt.Println(ptr)  // output: 0xc0000140a8
	fmt.Println(*ptr) // output: 10

	// But you do not have to deref struct to get it's values/methods
	type noDeref struct {
		value string
	}
	structPtr := &noDeref{value: "Value"}
	fmt.Printf(
		"At address %p we have %T struct with value %q\n",
		structPtr,
		structPtr,
		structPtr.value,
	)

	// Pass by value and by reference
	fmt.Print("Pass by value\n\n")
	addOne := func(x *int) {
		*x++
	}

	x = 10
	fmt.Println(x) // output: 10
	addOne(&x)
	fmt.Println(x) // output: 11
}

// Dynamic memory allocation is rarely used so we skip it here
// Pointers arithmetics is limited and could only be done with `unsafe` package
// ptr = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + uintptr(unsafe.Sizeof(arr[0]))))
// You probably do not need it

