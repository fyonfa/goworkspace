//selection-sort: it is a sort program created by me, just for fun and practice.
//time that i took to develop: start time:12.20h - end time:12.43(coffee included.) and debugging.
package main

import (
	"fmt"
)

//package scope, in order to use it in the for loop
var index int
var found bool

func main() {
	//arraySize will store the size of the array, max of 100 capacity
	var arraySize int

	//variable to store the smallest number
	var smallest int
	var tempValue int

	//Declaring array to be sorted
	//x := []int{8, 5, 2, 6, 9, 3, 1, 4, 0}

	fmt.Println("Please how many numbers will be inserted, max 100 numbers: ")
	fmt.Scanln(&arraySize)

	//declaring array of arraySize size
	x := make([]int, arraySize, 100)

	//length of the array
	l := len(x)

	//asking the user to input the numbers:
	fmt.Println("Please insert the", arraySize, "numbers, press enter after each number:")

	//declaring k and count for the input user loop
	k, count := 0, 0

	for count < arraySize {
		fmt.Scanln(&x[k])
		k++
		count++
	}

	//nested for loops to do the selection sort:
	for i := 0; i < l; i++ {
		smallest = x[i]
		found = false
		for j := i; j < l-1; j++ {
			if x[j+1] < smallest {
				smallest = x[j+1]
				index = j + 1
				found = true
			}
		}
		if found == false {
			continue
		}
		tempValue = x[i]
		x[i] = smallest
		x[index] = tempValue
		index = 0

	}

	fmt.Println(x)

}
