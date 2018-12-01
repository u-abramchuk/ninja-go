package main

import (
	"fmt"
)

const (
	c     = 13
	d int = 8
)

func main() {
	const a = 5
	const b int = 7

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)
}
