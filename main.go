package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, scala devs!")

	// PointersExample()
	// PointersPassByValueToCopy()
	PointersPassByValueToCopyFail()
	// TypeCompositionExample()
	IfsAndLoopsExample()
	// MapLoopExample()
	// ErrorHandlingExample()
	// DeferExample()
	// ArrayIndexesExample()
}

// Lack of imutability - constants are compile-time constants, and only primitive types. Read https://go.dev/doc/effective_go#constants
// Pointers, pass by value, pass by reference
// Type system: structs, receiver functions, interfaces, avoid interface pointers
// No monads and collection API - back to `if` and `for` loops
// No options, nil dereference
// Error handling, handle all errors
// Defer statement
// Anti-framework philisofy - use libs, built in and third-party
// Always read the docs: for example `time.Now()` returns local time. Use time.Now().UTC() or .UnixMili() for writing to DB
// Project structure, internal packages, forbidden cyclical imports
// Concurrency: everything is blocking by default, use goroutines, channels, wait gruops, mutexes and others from `sync` package https://gobyexample.com/goroutines
// Fun fact: you can state index of element when creating array
