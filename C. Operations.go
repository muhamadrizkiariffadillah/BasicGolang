package main

import "fmt"

func main() {
	//	Arithmetic Operator
	A, B := 23, 3
	var addition = A + B
	fmt.Println(addition)
	var subtraction = A - B
	fmt.Println(subtraction)
	var multiplication = A * B
	fmt.Println(multiplication)
	var division = A / B
	fmt.Println(division)
	var modulus = A % B
	fmt.Println(modulus)
	var (
		//	Comparison operators, the result will be true or false
		IsSame      = A == B
		IsNotSame   = A != B
		MoreThan    = A > B
		LessThan    = A < B
		MoreorEqual = A >= B
		LessOrEqual = A <= B
	)
	fmt.Println(IsSame)
	fmt.Println(IsNotSame)
	fmt.Println(MoreThan)
	fmt.Println(LessThan)
	fmt.Println(MoreorEqual)
	fmt.Println(LessOrEqual)
	//	Logical Operator
	var (
		left          = true
		right         = false
		and           = left && right
		or            = left || right
		negationTrue  = !left
		negationFalse = !right
	)
	fmt.Println(and, or, negationTrue, negationFalse)
}
