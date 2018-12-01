package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p *person) speak() {
	fmt.Println(p.first, p.last, p.age)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		first: "f",
		last:  "l",
		age:   2,
	}

	// saySomething(p)
	saySomething(&p)
}
