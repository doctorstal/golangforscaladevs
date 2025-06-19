package main

import (
	"fmt"
	"log"
)

func canProduceError(shouldError bool) (string, error) {
	if shouldError {
		return "", fmt.Errorf("here is your error")
	} else {
		return "Success", nil
	}
}

func ErrorHandlingExample() {
	res1, err := canProduceError(false)
	if err != nil {
		log.Fatal("Received an error!", err)
	}
	// err is not checked, but compilation does not fail. Take linter
	res2, err := canProduceError(true)

	fmt.Println(res1, res2)
}
