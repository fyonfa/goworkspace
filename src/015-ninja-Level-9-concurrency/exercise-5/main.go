package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius

}

func (s square) area() float64 {

	return s.length * s.length

}

type shape interface { //they implicit implement interface circle and square implemenat aread()
	//a value can be more than one type
	//any type who has this method area is also of type shape
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {

	circ := circle{ //this value circ is type circle, but circle has attached area() so also is type area() and area interface with shape, so also circ has type shape
		radius: 12.345,
	}
	squa := square{
		length: 15,
	}

	info(circ)
	info(squa)
}
