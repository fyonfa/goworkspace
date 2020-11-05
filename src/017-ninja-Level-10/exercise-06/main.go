package main

import (
	"fmt"
)

/***************************************************/
//write a program that puts 100 numbers to a channel/
//pull the numbers off the channel and print them	/
/***************************************************/
func main() {

	c := make(chan int)
	//puts the numbers into a channel
	c = putNumbers()

	//pull off numbers and display it
	pullNumbers(c)
	fmt.Println("about to exit")

}

func putNumbers() chan int {
	c := make(chan int) //

	//need a go routine who launches itself
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func pullNumbers(c chan int) {

	for v := range c {
		fmt.Println(v)
	}
}
