package main

import (
	"fmt"
)

func main() {
	foo(func(n interface{}) {
		fmt.Println(n)
	})
}

func foo(print func(interface{})) {
	print(42)
}
