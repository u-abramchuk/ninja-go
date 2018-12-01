package main

import (
	"fmt"
)

func main() {
	foo()
}

func foo() {
	defer fmt.Println("defer1")
	defer func() {
		fmt.Println("defer2")
	}()

	fmt.Println("println1")
	fmt.Println("println2")
}
