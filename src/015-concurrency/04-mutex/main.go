package main

import (
	"fmt"
	"runtime"
	"sync"
	//"time"
)

//this code has a race condition, multipe go routines accessing a shared variable
func main() {
	//we dont ave to set gomaxpro, by default 1
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter //v has this scope only so, every go fct will be accessing different v but counter will be the same memory area
			//time.Sleep(time.Second) //if you want to send it to sleep
			runtime.Gosched() //hey, go and run something else if you want
			v++
			counter = v
			mu.Unlock() //this code got lock down
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count: ", counter)
}
