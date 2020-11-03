package main

import (
	"fmt"
)

func main() {
	c := make(<-chan int, 2) // <- this means that only receive values to the channel, it depends when I put it, before or after chan

	//c <- 42 //In this case I am able to receive only in print, not to send it
	//c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("------")
	fmt.Printf("%T\n", c)

}
