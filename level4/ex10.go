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

	m["me"] = []string{`sleep`, `read`, `nothing`}

	for k, v := range m {
		fmt.Println(k)
		for i, val := range v {
			fmt.Printf("\t%v\t%v\n", i, val)
		}
	}

	fmt.Println("========================")

	delete(m, "no_dr")

	for k, v := range m {
		fmt.Println(k)
		for i, val := range v {
			fmt.Printf("\t%v\t%v\n", i, val)
		}
	}
}
