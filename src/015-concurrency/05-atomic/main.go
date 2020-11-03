package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	//"time"
)

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	var counter int64 //package atomic
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1) //write to my counter
			runtime.Gosched()
			fmt.Println("Counter\t", atomic.LoadInt64(&counter)) //and read from my counters
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count: ", counter)
}
