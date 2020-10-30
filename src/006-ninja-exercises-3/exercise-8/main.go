package main

import (
	"fmt"
)

func main() {
	switch { //by default is true when nothing is here
	case false:
		fmt.Println("should not print")
	case true:
		fmt.Println("should print")
	}
}
