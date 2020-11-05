package main

import (
	"fmt"
)

type person struct {
	name string
}

func main() {
	p1 := person{
		name: "James Bond",
	}

	fmt.Println(p1)
	fmt.Println(&p1.name)
	changeMe(&p1)
	fmt.Println(p1)

}

func changeMe(p *person) {

	fmt.Println("from changeme:", (*p).name)
	fmt.Println("from changeme:", &p.name)
	fmt.Println("from changeme:", p.name)
	p.name = "Miss money" //actually, this is what I am expecting: (*p).name = "xxx" but it can be done like that also I guess because it is already *person type
	fmt.Println("from changeme:", &p.name)

	fmt.Println("from changeme:", p.name) //but it can be done like this
	//(*p).name = "missp" //this is how it has to be don

}
