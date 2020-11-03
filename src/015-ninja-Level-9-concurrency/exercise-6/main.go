package main

import (
	"fmt"
)

func main() {
	//anonymous self executing function
	func() { //go func for concurrency if we have more than 1 cpu runs in parellel, we will see this in the next modules

		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}

	}() //for anonymous fct add this () here, self executing itself
	fmt.Println("done")

}
