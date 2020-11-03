package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	//send
	go foo(c)

	//receive, Go bar(c)
	bar(c) //so, if i remove the GO from the front of the function othis function will wait and we dont have to use a wait group, because it will block the channel until foo and bar are executed at the same time

	fmt.Println("about to exit")
}

//send
func foo(c chan<- int) { //send only channel, scope of c is only for this function
	c <- 42

}

//receive
func bar(c <-chan int) {
	fmt.Println(<-c)

}
