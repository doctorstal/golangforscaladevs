package main

import "fmt"

func ArrayIndexesExample() {
	names := []string{
		1: "alice",
		2: "bob",
	}
	for i, name := range names {
		fmt.Printf("Index %d, name %q\n", i, name)
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

var DayN = [...]string{
	"Unknown",
	"Tuesday",
	"Monday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
	"Sunday",
}

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

func IsWeekend(day Day) bool {
	return day == Sunday || day == Saturday
}
var fakeDay = 15
var isweekend = IsWeekend(Day(fakeDay))
