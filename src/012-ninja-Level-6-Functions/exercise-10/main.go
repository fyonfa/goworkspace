package main

import (
	"fmt"
)

//closure: enclosing scope of a variable in some code block
func main() {
	f := foo()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	g := foo() //enclosing in different scope, same x but different scope, that is why goes back to count to 1 to 5
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())

}

func foo() func() int { //return a func that returns an int
	x := 0              //this function is enclosing the scope of x
	return func() int { //returnning a func thar returns an int
		x++
		return x
	}
}
