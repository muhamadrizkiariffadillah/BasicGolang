package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	//	looping without argument
	count := 0
	for {
		fmt.Print(count, " ,")
		count++
		if count == 10 {
			break
		}
	}
	fmt.Println()
	//	range
	FullName := "Muhamad Rizki"
	for index, value := range FullName {
		fmt.Println("Index: ", index, "Value: ", value)
	}
	//	break and continue
	for x := 0; x < 10; x++ {
		if x == 5 {
			continue
		} else {
			if x == 9 {
				break
			}
		}
		fmt.Print(x)
	}
	fmt.Println()
}
