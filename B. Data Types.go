package main

import (
	"fmt"
	"reflect"
)

func main() {
	//	string
	var myName string = "Kibo"
	fmt.Println(reflect.TypeOf(myName), myName)
	//	Number
	// 	Int8,16,32,64 default depend on os bit
	var myAge int = 25
	fmt.Println(reflect.TypeOf(myAge), myAge)
	//	float 32,64
	var myTall float64 = 160.5
	fmt.Println(reflect.TypeOf(myTall), myTall)
	//	boolen true or false
	var handsome bool = true
	fmt.Println(reflect.TypeOf(handsome), handsome)
	//	number without negative value uint8,16,32,64
	var myNumber uint = 32
	fmt.Println(reflect.TypeOf(myNumber), myNumber)
	//	Array
	var MyArray = [5]int{1, 2, 3, 4, 5}
	fmt.Println(reflect.TypeOf(MyArray), MyArray)
	//	Slice without any limitation
	var MySlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(reflect.TypeOf(MySlice), MySlice)
	//	Map key and value
	var MyMap = map[int]string{
		1: "You can address me kibo.",
		2: "I'm cyber security engineer",
		3: "Learning everything to get knowledge about bug bounty",
		4: "I'll be a success person",
		5: "I hope. I can join with your organization.",
	}
	fmt.Println(reflect.TypeOf(MyMap), MyMap)
}
