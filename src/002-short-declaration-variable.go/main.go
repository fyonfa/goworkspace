package main

import "fmt"

var test = 43 //outside of a function body, and test is the scope of all

var z int //Variable with a identifier "z", ans it is type integer, Assigns the Zero value of type int to z. also I can do this, I have to say what type it is If I dont initialise it

func main() {

	//declaring and assigning a variable
	//short declaration operator
	x := 42 //this one only works between the main body, always short declaration operator inside of a body

	//what is the difference whit this declaration?
	var m = 43

	println(m)

	fmt.Println(x)
	x = 99
	fmt.Println(x)
	y := 100 + 70
	fmt.Println(y)
	z := "Bond, james"
	println(z)
	w := "hola" + z
	println(w)

	foo()

}

func foo() {
	fmt.Println("again: ", test)
}
