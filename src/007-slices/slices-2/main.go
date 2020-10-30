package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5}

	fmt.Println(x[:])   //take everything
	fmt.Println(x[1:])  // from position 1 to the end
	fmt.Println(x[1:3]) // from position 1 to 3 but not including 3
	fmt.Println(x[:3])  //from 0 up to but not including 3

	fmt.Println("appending:")

	myArray1 := []int{7, 8, 9, 10}
	myArray2 := []int{1, 2, 3}

	myArray2 = append(myArray2, 44, 55, 66)

	fmt.Println(myArray2)

	myArray1 = append(myArray1, myArray2...) //append everything in array2

	fmt.Println(myArray1)

	fmt.Println("I am deliting now:")

	myArray1 = append(myArray1[:2], myArray1[4:]...) //I took out the 10
	fmt.Println(myArray1)

}
