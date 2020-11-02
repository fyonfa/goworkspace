package main

import (
	"fmt"
)

func main() {
	//anonymous struct

	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "james",
		friends: map[string]int{
			"moneypenny": 555,
			"Q":          777,
			"M":          888,
		},
		favDrinks: []string{
			"martini",
			"water",
		},
	}
	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks)

	for k, v := range s.friends {
		fmt.Println(k, v)
	}
	for i, val := range s.favDrinks {
		fmt.Println(i, val)
	}

}
