package main

import "fmt"

func main() {
	var full_name string = "Muhamad Rizki Arif Fadillah"
	var age int = 25
	var tall float64 = 160.5
	var handsome bool = true
	fmt.Println(full_name, age, tall, handsome)
	//	Public Variable and Camel case
	var FullName = "Muhamad Rizki Arif Fadillah"
	var MyAge = 25
	var MyTall = 160.5
	var AmIHandsome = true
	fmt.Println(FullName, MyAge, MyTall, AmIHandsome)
	YourName := "Muhamad Rizki Arif Fadillah"
	YourAge := 25
	YourTall := 160.5
	AreYouHandsome := true
	fmt.Println(YourName, YourAge, YourTall, AreYouHandsome)
	//	Constanta can't be changed.
	const MyReallyName = "Muhamad Rizki"
	fmt.Println(MyReallyName)
}
