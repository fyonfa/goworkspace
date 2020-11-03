package main

import (
	"fmt"
)

//recursion

//{\displaystyle 0,\;1,\;1,\;2,\;3,\;5,\;8,\;13,\;21,\;34,\;55,\;89,\;144,\;\ldots }

func main() {

	i := 45
	fmt.Printf("The fibbonacci sequence of %d is %d", i, fibbonacci(i))

}

func fibbonacci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibbonacci(i-1) + fibbonacci(i-2)
}
