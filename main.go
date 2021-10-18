package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("My string is set to", myString)
	changeUsingPointer(&myString)
	log.Println("after func call My string is set to", myString)
}

func changeUsingPointer(s *string) {
	newValue := "Red"
	*s = newValue
}
