package main

import (
	"fmt"
	"runtime"
	"sync"
)

//func init() { //initialization, to set things up

var wg sync.WaitGroup //wg is a type from waitgroup from package sync, wg has package scope.

func main() { //this is a main Go routine, main has already one
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GO routines\t", runtime.NumGoroutine())

	wg.Add(1) //1.-one thing that we are going to be waiting for, so add that to wg, after this we launching foo()
	go foo()  //so, main is launched, and at this point we lauch foo, other routine, and it continues in parallel, and qhen main finished, everything finishes
	//we need some sort of synchronisation... "PRIMITIVE SYNCRONISATION-CONCURRENCY"
	bar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GO routines\t", runtime.NumGoroutine())
	wg.Wait() //3.-and here it is going to wait until all the thing added are done
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done() //2.- when foo does what it needs to do, it is done, we say done, and it will go to wait, this done will remove the thing that we added in wg.add
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
