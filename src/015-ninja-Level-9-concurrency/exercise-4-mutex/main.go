package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*****Check this exercise as is given me 1 or 2 races when I am doing go run -race main.go****/
func main() {
	var wg sync.WaitGroup

	incrementer := 0
	gs := 100
	wg.Add(gs)
	var m sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			fmt.Println(incrementer)
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", incrementer)
}

/*
package main

import (
	"fmt"
	"sync"
)

/****************Method Sets***************
//fixing the problem from before using mutex, mutual exclusion lock

func main() {
	var wg sync.WaitGroup

	incrementer := 0

	gs := 100
	wg.Add(gs)
	var m sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {

			m.Lock()
			v := incrementer
			v++
			incrementer = v
			fmt.Println(incrementer)
			m.Unlock()
			wg.Done()

			/*
				v := incrementer //it make sense to remove runtime.Gosched
				runtime.Gosched()
				v++
				incrementer = v
				fmt.Println(incrementer)
				wg.Done()
			*
		}()
	}
	wg.Wait()
	fmt.Println(incrementer)

}

*/
