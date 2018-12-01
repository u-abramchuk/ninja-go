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

	fmt.Println(person1.firstName)
	fmt.Println(person1.lastName)

	for _, v := range person1.favoriteIceCreamFlavors {
		fmt.Println(v)
	}

	fmt.Println(person2.firstName)
	fmt.Println(person2.lastName)

	for _, v := range person2.favoriteIceCreamFlavors {
		fmt.Println(v)
	}

}
