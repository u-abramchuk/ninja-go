package main

import (
	"fmt"
)

type Person struct {
	firstName               string
	lastName                string
	favoriteIceCreamFlavors []string
}

func main() {
	person1 := Person{
		firstName:               "John",
		lastName:                "Smith",
		favoriteIceCreamFlavors: []string{"Vanilla"},
	}

	person2 := Person{
		firstName:               "Mary",
		lastName:                "Johns",
		favoriteIceCreamFlavors: []string{"Strawberry"},
	}

	m := map[string]Person{
		person1.lastName: person1,
		person2.lastName: person2,
	}

	for k, v := range m {
		fmt.Println(v.firstName)
		fmt.Println(k)

		for _, v := range v.favoriteIceCreamFlavors {
			fmt.Println(v)
		}
	}
}
