/******************************************************/
//write a program that launches 10 numbers to a channel/
//		each goroutine adds 10 numbers to a channel	   /
//pull the numbers off the channel and print them	   /
/******************************************************/

package main

import "fmt"

func main() {
	c := make(chan int)

	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c)
	}

	fmt.Println("about to exit")
}

/*
package main

import (
	"fmt"
	"runtime"
)

func main() {
	x := 10
	y := 10
	c := gen(x, y)

	for i := 0; i < x*y; i++ {
		fmt.Println(i, <-c)
	}
	fmt.Println("ROUTINES", runtime.NumGoroutine())
}

func gen(x, y int) <-chan int {
	c := make(chan int)

	for i := 0; i < x; i++ {
		go func() {
			for j := 0; j < y; j++ {
				c <- j
			}
		}()
		fmt.Println("ROUTINES", runtime.NumGoroutine())
	}
	return c
}

/*
func main() {
	c := make(chan int)

	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			//close(c)//if I close c here it is not going to work because I wanto close c once, not 100 times.
		}()
		//close here neither
	}
	//neither here if i want to add values

	for k := 0; k < 100; k++ { //we know exactly how many values we have
		fmt.Println(<-c)
	}

}
*/
