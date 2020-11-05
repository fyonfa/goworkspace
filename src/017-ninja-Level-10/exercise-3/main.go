package main

import (
	"fmt"
)

//check video 172 for epxlanation
func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive(c <-chan int) {

	for v := range c { //this is going to block until c is closed
		fmt.Println(v)
	}

}
