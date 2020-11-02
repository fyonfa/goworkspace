package main

import (
	"fmt"
)

func main() {

	m := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	//fmt.Println(m)
	m["Flemming_ian"] = []string{`steaks`, "cigars", "espionage"} //adding one more element
	//delete a record
	delete(m, "no_dr")
	for k, v := range m {
		fmt.Println("This is sthe record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

}
