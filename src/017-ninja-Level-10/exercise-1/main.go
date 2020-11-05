package main

import (
	"fmt"
)

func main() {
	c := make(chan int) // I can add one here, creating one buffer or also a self executing anonymous fct
	go func() {
		c <- 42
	}()

	fmt.Println(<-c) //this will block until go routine executes and when both are executed at tthe same time, this will receive.
}
