package main

import (
	"fmt"
)

func MapLoopExample() {
	kvMap := make(map[string]string)

	// golang:
	var mymap map[string]map[string]string
	// Scala: Map[String, Map[String, String]
	kvMap["key"] = "value"
	for k, v := range kvMap {
		fmt.Printf("%q:%q\n", k, v)
	}

	kvMap2 := map[string]string{
		"key": "value",
		"key2": "value",
	}
	for k, v := range kvMap2 {
		fmt.Printf("%q:%q\n", k, v)
	}

	fmt.Printf("zero value: %q\n", kvMap["a"])


	delete(kvMap, "key")

	if v, ok := kvMap["unknown"]; !ok {
		fmt.Println("No such key")
	} else {
		fmt.Printf("Found value %q\n", v)
	}

	mapdef := mapwithdefault(kvMap)
	defval := mapdef.getOrDefault("unknown", "default")

	fmt.Printf("Got from map: %q\n", defval)

	defmap := defaultmap{kvMap, "default"}
	defval = defmap.get("unknown")

	fmt.Printf("Got from defmap: %q\n", defval)
}

type mapwithdefault map[string]string

func (m mapwithdefault) getOrDefault(key, def string) string {
	if v, ok := m[key]; ok {
		return v
	} else {
		return def
	}
}

type defaultmap struct {
	mapwithdefault
	def string
}

func (d *defaultmap) get(key string) string {
	return d.getOrDefault(key, d.def)
}
