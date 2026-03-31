package main

import "fmt"

func mightPanic(input int) int {
	if input == 0 {
		panic("cannot divide by zero")
	}
	return 100 / input
}

// safeDiv wraps a panicking call and recovers into an error.
// recover() only works when called directly inside a deferred function.
func safeDiv(input int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from panic: %v", r)
		}
	}()
	result = mightPanic(input)
	return
}

func PanicAndRecoveryExample() {
	// Normal case - no panic
	if res, err := safeDiv(5); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", res)
	}

	// Panicking case - recovered into an error
	if res, err := safeDiv(0); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", res)
	}
}
