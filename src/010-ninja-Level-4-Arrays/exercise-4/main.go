package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55) //append has variatic , unnumber of elements to append
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}

	x = append(x, y...) //to append all the y[]
	fmt.Println(x)

}
