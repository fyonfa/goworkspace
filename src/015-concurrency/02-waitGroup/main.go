package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2) //waiting for 2 things
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	fmt.Println("This is fct foo")
	wg.Done()
}

func bar() {
	fmt.Println("This is fct bar")
	wg.Done()
}
