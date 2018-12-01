package main

import (
	"fmt"
)

type person struct {
	name string
}

func changeMe(p *person) {
	p.name = "changed"
}

func main() {
	p := person{
		name: "John",
	}

	fmt.Println(p)

	changeMe(&p)

	fmt.Println(p)
}
