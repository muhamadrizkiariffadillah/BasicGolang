package main

import "fmt"
func main() {
	myName := TypeString().(string)
	fmt.Println(myName)
	myAge := TypeInt().(int)
	myTall := TypeFloat().(float64)
	isHandsome := TypeBoolen().(bool)
	fmt.Println(myAge,myTall,isHandsome)
}

func TypeString()any{
	return "My name is Muhamad Rizki"
}
func TypeInt()any{
	return 25
}
func TypeFloat()any{
	return 169.9
}
func TypeBoolen()any{
	return true
}
