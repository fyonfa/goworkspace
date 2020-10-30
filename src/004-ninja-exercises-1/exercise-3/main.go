package main

import "fmt"

var x = 42
var y = "James Bond"
var z = true

func main() {

	//s := fmt.Sprintf("%d\n%s\n%t", x, y, z) //also like below, v is default format
	s := fmt.Sprintf("%v\n%v\n%v", x, y, z)
	fmt.Println(s)

}
