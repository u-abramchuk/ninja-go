package main

import (
	"fmt"
)

func main() {
	nxt1 := next(1)
	fmt.Println(nxt1())
	fmt.Println(nxt1())
	fmt.Println(nxt1())

	nxt2 := next(5)
	fmt.Println(nxt2())
	fmt.Println(nxt2())
	fmt.Println(nxt2())
	fmt.Println(nxt2())
}

func next(n int) func() int {
	x := n
	return func() int {
		x *= 2

		return x
	}
}
