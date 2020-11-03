package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

//func (r receiver) identifier(parameters) return(s) {code}

func (p person) speak() { //function is speak and we are going to attach the fct speak to the p person parameter, any value of this type has acces to this method
	//is going to attach this fct to that type when we have a receiver
	fmt.Println("I am,", p.first, p.last, "and I am", p.age, "years old.")
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	p1.speak()
}
