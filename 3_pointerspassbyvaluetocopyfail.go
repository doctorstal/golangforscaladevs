package main

import (
	"fmt"
	"strings"
)

func PointersPassByValueToCopyFail() {
	fmt.Printf("Copy with pointers is not so simple:\n")

	type person struct {
		Name    string
		Surname string
		Kids    []*person
	}
	renamePersonAndKids := func(p person, name, surname string) person {
		p.Name = name
		p.Surname = surname
		for _, kid := range p.Kids {
			kid.Surname = surname
		}
		return p
	}
	var printPerson func(*person) string
	printPerson = func(p *person) string {
		var builder strings.Builder
		builder.WriteString(fmt.Sprintf("Name: %v, Surname %v\n", p.Name, p.Surname))
		for _, kid := range p.Kids {
			builder.WriteString(printPerson(kid))
		}
		return builder.String()
	}

	aliceSurname := "Koval"
	aliceWithKids := person{Name: "Alice", Surname: aliceSurname, Kids: []*person{
		{
			Name:    "Jane",
			Surname: aliceSurname,
		},
	}}
	fmt.Printf("alice %v\n", printPerson(&aliceWithKids))

	bobWithKids := renamePersonAndKids(aliceWithKids, "Bob", "Melnyk")

	fmt.Printf("alice: %v\nbob: %v \n", printPerson(&aliceWithKids), printPerson(&bobWithKids))
}
