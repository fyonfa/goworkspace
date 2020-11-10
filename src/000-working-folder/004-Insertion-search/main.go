//selection-sort: it is a sort program created by me, just for fun and practice.
//time that i took to develop: start time:12.20h - end time:12.43(coffee included.) and debugging.
package main

import (
	"fmt"
)

var index int
var temp int

func main() {
	//declaring array of values
	x := []int{5, 4, 3, 2, 1}

	//getting the length of the array
	l := len(x)
	var temp int

	for i := 1; i < l; i++ {
		temp = x[i]

		j := i - 1
		for j > -1 && temp < x[j] {
			x[j+1] = x[j]
			j--
		}
		x[j+1] = temp

	}

	fmt.Println(x)

}
