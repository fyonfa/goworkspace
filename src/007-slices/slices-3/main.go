package main

import (
	"fmt"
)

func main() {

	/*****MAKE******/
	/*Usually as we now from C or C++ we declare an array and we have to destroy, and create an array of bigger size in order to add more items to the existing array... sort of
	  with slicing, the slice is sitting on top of other internal array which gets destroyed and created dependingof the slice of the user, so that is what happens behid the scenes
	  with MAKE, if you already know the size of your array, you declare this slice/array and it wont do that destroying and creating in the background, saving time to the compiler
	  and saving executing time and ganing efficiency...*/

	x := make([]int, 10, 12) //it needs the type, the length and the capacity

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 42
	x[9] = 999

	//x[10] = 100 I can´t do this because it is out of our slice
	//but I can do this however because the lenght size is 100 I guess:
	x = append(x, 1234)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 1235)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 1236) //here it exceeds the capacity of the array, so it doubles it, so it gets an underlying array, better use it just when you know the size and capacity that you are going to use.

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}
