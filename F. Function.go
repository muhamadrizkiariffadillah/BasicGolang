package main

import (
	"fmt"
	"strings"
)

func main() {
	HelloWorld()
	helloWorld()

	saveReturn := ReturnHelloWorld()
	fmt.Println(saveReturn)

	saveParameter := ParameterFunction("Muhamad Rizki ", "Arif Fadillah")
	fmt.Println(saveParameter)

	saveMultiReturn, saveMultiReturn2 := MultipleReturn("Muhamada Rizki Arif Fadillah")
	fmt.Printf("Name: %s, the length of name: %d\n", saveMultiReturn, saveMultiReturn2)

	variadic := variadicFunction(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
	fmt.Println("Average :", variadic)

	// Closure
	anonymousFunction := func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				min, max = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}
	number := []int{1, 2, 9, 8, 3, 5, 6, 7, 8, 3, 11, 19, 28, 57}
	min, max := anonymousFunction(number)
	fmt.Printf("Minimum: %d, Maximum: %d\n", min, max)
	// Immediately-Invoked Function Expression (IIFE)
	var numbers = []int{1, 9, 8, 7, 5, 4, 6, 3, 2, 1, 4, 5, 6}
	var getMinMax = func(min int) []int {
		var r []int
		for _, value := range numbers {
			if value < min {
				continue
			}
			r = append(r, value)
		}
		return r
	}(3)
	fmt.Println(getMinMax)
	// Function as parameter
	var data = []string{"Rizki","Rangga","Rahman","Kibo"}
	var dataContainA = filter(data,func(s string) bool {
		return strings.Contains(s,"a")
	})
	var dataContainI = filter(data,func(s string) bool {
		return strings.Contains(s,"i")
	})
	fmt.Println("Filter Contain A:",dataContainA)
	fmt.Println("Filter Contain I:",dataContainI)


}

// Public function
func HelloWorld() {
	fmt.Println("Hello World!")
}

// private function
func helloWorld() {
	fmt.Println("Hello World")
}

// Return value
func ReturnHelloWorld() string {
	return "Hello World!"
}

// Parameter function
func ParameterFunction(firstName, lastName string) string {
	fullName := firstName + lastName
	return fullName
}

// Multiple return function
func MultipleReturn(fullname string) (string, int) {
	return fullname, len(fullname)
}

// Variadic function
func variadicFunction(numbers ...int) float64 {
	total := 0
	for _, number := range numbers {
		total += number
	}
	average := float64(total) / float64(len(numbers))
	return average
}

// Function as parameter
func filter(data []string, calback func(string) bool) []string {
	result := []string{}
	for _, value := range data {
		if filtered := calback(value); filtered {
			result = append(result, value)
		}
	}
	return result
}
