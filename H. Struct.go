package main

import (
	"fmt"
)

// Struct or structure
// It's like an object in others programming language

type Person struct {
	id uint
	name string
	age uint
}

func main() {
	firstPerson := Person{
		id: 1,
		name: "Rizki",
		age: 20,
	}
	fmt.Println(firstPerson)
	// Pointer object variable

	var s1 = Person{
		1,
		"Muhamad Rizki",
		20,
	}

	var s2 *Person = &s1
	fmt.Println(s1)
	fmt.Println(s2)

	s2.name = "Kibo"

	fmt.Println(s1)
	// Anonymous struct
	anon := struct{
		id int
		fullname string

	}{
		id: 1,
		fullname: "Kibo",
	}
	fmt.Println(anon)
	// slice anonymous
	students := []Person{
		{
			id: 1,
			name: "Kibo",
			age: 20,
		},
		{
			id: 2,
			name: "Arif",
			age: 30,
		},
		{
			id: 3,
			name: "Susu",
			age: 30,
		},
	}
	// Iteration using looping
	for _,y := range students{
		fmt.Println(y)

	}
}
