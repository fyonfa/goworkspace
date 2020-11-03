package main

import (
	"fmt"
)

//so, starts in package main, go to fct main, create a channel, tries to put 42 on the channel and crash.when you try to send and receive from a channel,
//the transaction cannot occur if receive and send doesnt happen at the same time. Send and receive at the same time. And if they can´t happen at the same
//time it blocks the send and receiv blocks until the receiver is ready to pull it off.
//CHANNELS BLOCK

func main() {
	c := make(chan int)

	c <- 42

	fmt.Println(<-c)
}
