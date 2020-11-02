package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("hello play") //with defer it will skip the fct and do the print first and it will caome back later on, at the end
}

func foo() {
	defer func() {
		fmt.Println("foo defer ran")
	}() //anonymous fct
	fmt.Println("foo ran")
}
