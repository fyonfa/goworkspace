package main

import (
	"fmt"
)

func main() {
	//having a key of string and slice of string as values
	//have a warning with this
	/*	m := map[string][]string{
			`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
			`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
			`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
		}
	*/
	m := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	//fmt.Println(m)

	for k, v := range m {
		fmt.Println("This is sthe record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
