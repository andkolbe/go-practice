package main

import (
	"fmt"

	"github.com/andkolbe/go-practice/helpers"
)

func main() {
	var myVar helpers.SomeType
	myVar.TypeName = "Some name"
	fmt.Println(myVar.TypeName)
}
