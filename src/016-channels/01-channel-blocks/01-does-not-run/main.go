package main

import (
	"fmt"
)

//so, starts in package main, go to fct main, create a channel, tries to put 42 on the channel and crash.when you try to send and receive from a channel,
//the transaction cannot occur if receive and send doesnt happen at the same time. Send and receive at the same time. And if they can´t happen at the same
//time it blocks the send and receiv blocks until the receiver is ready to pull it off.
//CHANNELS BLOCK

func main() { //here difference between first example -->*****try to stay away from buffer channels***
	c := make(chan int, 1) //++++++here is a channel where I can put integers, "1" this is a buffer channel, it will allow 1 to sit in there, one value to sit in there
	c <- 42                //put 42 on c, because there is room for it
	//c <- 43                //---->if I add this, it wont run, I would have to add 2 in make
	fmt.Println(<-c)
}
