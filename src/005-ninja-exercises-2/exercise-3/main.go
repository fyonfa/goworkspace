package main

import (
	"fmt"
)

const (
	a     = 42 //untyped constant of a type, give our compiler a bit more flexibility to decide what type it is, kind of auto I guess
	b int = 43 //typed constant
)

func main() {
	fmt.Println(a, b)
}
