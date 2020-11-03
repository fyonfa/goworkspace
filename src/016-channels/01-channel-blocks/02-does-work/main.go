package main

import (
	"fmt"
)

//so, starts in package main, go to fct main, create a channel, tries to put 42 on the channel and crash.when you try to send and receive from a channel,
//the transaction cannot occur if receive and send doesnt happen at the same time. Send and receive at the same time. And if they can´t happen at the same
//time it blocks the send and receiv blocks until the receiver is ready to pull it off.
//CHANNELS BLOCK

func main() {
	c := make(chan int) //here is a channel where I can put integers
	go func() {         //this will make it run, this launches it, the func it is running by its own and the flow continues
		c <- 42 //ready to put it on
	}()
	//put 42 on c
	fmt.Println(<-c) //and this one blocks until it gets the value off, ready to take it off
	//and we have coordination with concurrent design
}
