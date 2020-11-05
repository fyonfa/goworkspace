package main

import (
	"fmt"
)

func main() {
	n, err := fmt.Println("hello") //give me 6 because the new line char
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(n)
}
