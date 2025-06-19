package main

import (
	"fmt"
)

func MapLoopExample() {
	kvMap := make(map[string]string)
	kvMap["key"] = "value"
	kvMap["key2"] = "value2"
	for k, v := range kvMap {
		fmt.Printf("%q:%q\n", k, v)
	}
	fmt.Println()

	kvMap2 := map[string]string{
		"key": "value",
	}
	for k, v := range kvMap2 {
		fmt.Printf("%q:%q\n", k, v)
	}
	fmt.Println()

	fmt.Printf("zero value: %q\n", kvMap["a"])

	fmt.Printf("Map length: %d\n", len(kvMap2))
	delete(kvMap2, "key")
	fmt.Printf("Map length: %d\n", len(kvMap2))
	fmt.Println()

	if v, ok := kvMap["unknown"]; !ok {
		fmt.Println("No such key")
	} else {
		fmt.Printf("Found value %q\n", v)
	}
	fmt.Println()

	defMap := mapwithdefault(kvMap)
	fmt.Printf("unknown: %q\n", defMap.getOrDefault("unknown", "default value"))
	fmt.Printf("key %q\n", defMap.getOrDefault("key", "default value"))
	fmt.Println()

	defMap2 := defaultmap{kvMap, "default2"}

	fmt.Printf("unknown: %q\n", defMap2.get("unknown"))
	fmt.Printf("key %q\n", defMap2.get("key"))
}

type mapwithdefault map[string]string

func (m mapwithdefault) getOrDefault(key, def string) string {
	if v, ok := m[key]; !ok {
		return def
	} else {
		return v
	}
}

type defaultmap struct {
	mapwithdefault
	Default string
}

func (d *defaultmap) get(key string) string {
	return d.getOrDefault(key, d.Default)
}
