package main

import (
	"fmt"
)

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		favFlavors: []string{
			"Chocolate",
			"martini",
			"coke"},
	}

	p2 := person{
		first: "Miss",
		last:  "MoneyPenny",
		favFlavors: []string{
			"strawberryy",
			"vanilla",
			"capuchino",
		},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v) //key and value
	}
	fmt.Println()
	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favFlavors {
			fmt.Println(i, val)
		}
	}

}
