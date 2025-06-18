package main

import (
	"fmt"
)

func MapLoopExample() {
	kvMap := make(map[string]string)
	kvMap["key"] = "value"
	for k, v := range kvMap {
		fmt.Println(k, v)
	}

	kvMap2 := map[string]string{
		"key": "value",
	}
	for k, v := range kvMap2 {
		fmt.Println(k, v)
	}

	if v, ok := kvMap["unknown"]; !ok {
		fmt.Println("No such key")
	} else {
		fmt.Println("Found value", v)
	}

	fmt.Println("zero value:", kvMap["a"])

	delete(kvMap, "key")
}

