package main

import (
	"fmt"
)

func main() {
	f := foo()
	fmt.Println(f())

}

func foo() func() int { //function foo which return a function that returns an int
	//this return a function
	return func() int { //func that returns an int
		return 42
	}
}
