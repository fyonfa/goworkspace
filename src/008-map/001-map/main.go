package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James":           32,
		"miss moneypenny": 27,
		"Fernando":        34,
		"kath":            31,
	}
	fmt.Println(m)

	fmt.Println(m["James"])
	fmt.Println(m["miss moneypenny"])

	v, ok := m["James"] //idiomatic statement in Go that we will se a lot
	//v for value, ok for condition
	fmt.Println(v)
	fmt.Println(ok)

	v, ok = m["no name"] //idiomatic statement in Go that we will se a lot
	//v for value, ok for condition
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["miss moneypenny"]; ok {
		fmt.Println("This is the if print", v)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("for loop:")
	for i, v := range m {
		fmt.Println(i, v)
	}

	///delete a map key

	delete(m, "James")
	fmt.Println(m)

	if v, ok := m["miss moneypenny"]; ok {
		fmt.Println("This is the if print", v)
		delete(m, "miss moneypenny")
	}

}
