package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var random = rand.Int64N(20)
	fmt.Println(random)
	// if else
	if random > 0 && random < 10 {
		fmt.Println("Less than 10")
	} else if random >= 10 && random < 15 {
		fmt.Println("Less than 15")
	} else {
		fmt.Println("Less than or equals 20")
	}
	//	switch case
	switch random {
	case 1, 2, 3, 4, 5, 6, 7, 9:
		fmt.Println("C")

	case 10, 11, 12, 13, 14, 15:
		fmt.Println("B")
	default:
		fmt.Println("A")
	}
}
