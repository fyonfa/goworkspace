package main

import (
	"fmt"
)

func main() { //read from left to right
	c := make(chan<- int, 2) // <- this means that onlye send values to the channel

	c <- 42 //here I am able to send, but below, in print, I get error because I want to receive it
	c <- 43

	//fmt.Println(<-c)
	//fmt.Println(<-c)
	fmt.Println("------")
	fmt.Printf("%T\n", c)

}
