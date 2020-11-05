//Use this program to sort integers inserted from the user.
package main

import (
	"fmt"
)

//declare this variable scope to be able to change it in the for loop, it would be enclosed if not.
var found = true

func main() {

	var arraySize int
	//take array size input form the user
	fmt.Println("Please how many numbers will be inserted, max 100: ")
	fmt.Scanln(&arraySize)

	//make array maxim 100 capacity and size "arraySize
	x := make([]int, arraySize, 100)

	fmt.Println("Please insert the", arraySize, "numbers:")
	//declare j anc count
	j, count := 0, 0
	for count < arraySize {

		fmt.Scanln(&x[j])
		j++
		count++
	}

	//temporal variable to store int
	tempValue := 0
	l := len(x)

	//compare the first letter in the array with the next and if is higher swapped it and make "found" = true.

	for found == true {
		found = false
		for i := 0; i < l-1; i++ {
			if x[i] > x[i+1] {
				tempValue = x[i]
				x[i] = x[i+1]
				x[i+1] = tempValue
				found = true
			}

		}
	}

	//display the array
	fmt.Println(x)

}
