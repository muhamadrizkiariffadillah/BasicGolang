package main

import "fmt"

type Address struct {
	city,province,county string
}

func main() {
	address1 := Address{"Bandung","West Java","Indonesia"}
	address2 := &address1
	fmt.Println(address1)
	fmt.Println(address2)
	address2.city = "Garut"
	fmt.Println(address1)
	fmt.Println(address2)
	// New memory address
	// address2 = &Address{"Jakarta","DKI Jakarta","Indonesia"}
	// fmt.Println(address1)
	// fmt.Println(address2)
	// Same memory address
	*address2 = Address{"Purwakarta","West Java","Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}