package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	/*
		go func() {
			fmt.Println("Hi from one")
		}()

		go func() {
			fmt.Println("Hi from two")
		}()
		fmt.Println("about to exit")
	*/

	fmt.Println("begin Go Routines", runtime.NumGoroutine())
	fmt.Println("Begin CPUs", runtime.NumCPU())
	wg.Add(2) //waiting for 2 things
	go foo()
	go bar()
	fmt.Println("Mid Go routines", runtime.NumGoroutine())
	fmt.Println("Mid CPUs", runtime.NumCPU())
	wg.Wait()
	fmt.Println("About to exit")
	fmt.Println("End Go routines", runtime.NumGoroutine())
	fmt.Println("End CPUs", runtime.NumCPU())
}

func foo() {
	fmt.Println("This is fct foo")
	wg.Done()
}

func bar() {
	fmt.Println("This is fct bar")
	wg.Done()

}
