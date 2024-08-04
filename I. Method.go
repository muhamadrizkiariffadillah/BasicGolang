package main

import "fmt"

type Student struct {
	Id uint
	Name string
}

func main() {
	person := Student{
		Id: 1,
		Name: "Kibo",
	}
	fmt.Println(person.getName())
}

func (s Student) getName()(string)  {
	return s.Name
}

