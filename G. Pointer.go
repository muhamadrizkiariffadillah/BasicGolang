package main

import "fmt"

func main() {
	var person string = "Muhamad Rizki Arif Fadillah"
	var address *string = &person
	fmt.Println("Person Value: ",*address)
	fmt.Println("Person memory address: ",address)
}