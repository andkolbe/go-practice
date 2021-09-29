package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)
	changeUsingPointer(&myString)
	log.Println("after func call myString is set to", myString)

}

func changeUsingPointer(s *string) {
	newValue := "Red"
	*s = newValue
}

// a pointer points to a specific location in memory and gives you a means of getting that location in memory
// when you create a variable, you are creating a variable in memory
// you store a value at a particular memory location
