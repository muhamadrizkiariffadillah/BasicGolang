package main

import "fmt"

func main() {
	runApplication(true)
}

func logging(){
	fmt.Println("Monitoring")
	// Recover is a safety way to keep application running
	
	message := recover()

	fmt.Println("Error message: ",message)
}
func runApplication(error bool){
	defer logging()
	fmt.Println("Application is running")
	if error{
		// Panic is like raise in python
		panic("You raise the error")
	}
}