package main

import "fmt"

//create a type person struct
type person struct {
	Name     string
	Lastname string
	Age      int
}

//attach a method speak to type person using a pointer receiver
func (p *person) speak() {
	fmt.Println("hola, this is my full name:", p.Name, p.Lastname, ".", "And, my age:", p.Age)
	fmt.Printf("hola, this is my full name: %s %s, and my age %d", p.Name, p.Lastname, p.Age)
	fmt.Println()
}

//create a type human interface
//to implicitly implement the interface, a human must have the speak method
type human interface {
	speak()
}

func saySomething(h human) {

	h.speak()

}

func main() {

	p1 := person{
		Name:     "Fernando",
		Lastname: "Yonfa",
		Age:      34,
	}

	saySomething(&p1) //work with this
	p1.speak()        //also with this :)

}
