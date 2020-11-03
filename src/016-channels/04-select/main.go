package main

import (
	"fmt"
)

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//send
	go send(eve, odd, quit)

	//receive, pull values off of the channel
	receive(eve, odd, quit)

	fmt.Println("about to exit")

}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("from eve channel: ", v)
		case v := <-o:
			fmt.Println("from odd channel: ", v)
		case v := <-q:
			fmt.Println("from quit channel: ", v)

			return
		}

	}

}

func send(e, o, q chan<- int) { //func that sends values
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	//close(e)
	//close(o)
	//q <- 0
	close(q)
}
