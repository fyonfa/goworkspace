package main

import "fmt"

func main() {

	n := 0 //keep you scope as narrow as possible
	fmt.Println("Please Insert a number:")
	fmt.Scanf("%d", &n)

	fmt.Printf("%d\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%#X\n", n)

}
