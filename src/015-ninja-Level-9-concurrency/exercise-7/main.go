package main

import (
	"fmt"
)

var x int = 7
var g func() = func() { //we can do this also, but better var g func() as usual
	fmt.Println("G from outside")
} //var g is the type func()

//assign a func to a variable
func main() {

	f := func() {

		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}

	}
	f()
	fmt.Printf("%T\n", f)

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	g()
	g = f
	g()
	fmt.Printf("This is g %T\n", g)
	fmt.Println("done")
}
