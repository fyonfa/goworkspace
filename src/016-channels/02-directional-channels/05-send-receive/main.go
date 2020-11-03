package main

import (
	"fmt"
)

func main() {

	c := make(chan int, 2)
	cr := make(<-chan int) //receive
	cs := make(chan<- int) //send

	fmt.Println("-----")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("c\t%T\n", cr)
	fmt.Printf("c\t%T\n", cs)

	//specific to general doesn´t assign

	//c = cr
	//c = cs
	//specific to specific doesn´t assign, two different types
	//cs=cr

	//general to specific does assign

	//cr = c
	//cs = c
	//general to specific converts
	fmt.Println("-----")
	fmt.Printf("c\t%T\n", (<-chan int)(c)) //if I write here (cs) doesnt convert
	fmt.Printf("c\t%T\n", (chan<- int)(c))

}
