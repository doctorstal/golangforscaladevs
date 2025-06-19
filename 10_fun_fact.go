package main

import "fmt"

func ArrayIndexesExample() {
	names := []string{
		1: "alice",
		2: "bob",
	}

	for i, name := range names {
		fmt.Printf("Name at %d is %q\n", i, name)
	}
}

type Day int

const (
	Unknown Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

var DayNames = [...]string{
	Unknown:   "Unknown",
	Monday:    "Monday",
	Tuesday:   "Tuesday",
	Wednesday: "Wednesday",
	Thursday:  "Thursday",
	Friday:    "Friday",
	Saturday:  "Saturday",
	Sunday:    "Sunday",
}
