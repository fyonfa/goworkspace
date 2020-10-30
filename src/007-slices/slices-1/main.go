package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := []int{1, 2, 3, 4}

	//this is a standar loop
	for i := 0; i <= 3; i++ {
		fmt.Println(x[i])
		fmt.Println(reflect.TypeOf(x[i]))

	}

	fmt.Println(x)
	fmt.Println(reflect.TypeOf(x))

	//this is a better way to do it in go, similar to phyton i guess

	for i, v := range x {
		fmt.Println(i, v)

	}

}
