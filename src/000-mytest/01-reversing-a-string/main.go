package main

import (
	"fmt"
)

func main() {

	s := "hola esto va a ser revertido"
	c := make([]byte, len(s), 100)

	c[0] = s[2]

	for i, j := 0, len(s)-1; i <= len(s)-1; i, j = i+1, j-1 {
		c[i] = s[j]

	}

	fmt.Printf("%T, %T", s, c)
	fmt.Println(len(s), len(c))

	fmt.Println(s)
	fmt.Println(string(c))

}
