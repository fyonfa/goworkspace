package main

import "fmt"

var y int = 42

//declaring a variable with identifier z, is a string type and I assign value of " ", It is a static, not a dynamic language, variable is declared to a value of certain type
var z = "Shaken, not stritted"

// also I can do this: var z string = "Shaken, not stritted"

func foo() {
	fmt.Println("hi, this is Fct: foo")

	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)

}
