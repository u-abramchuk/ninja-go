package main

import (
	"fmt"
)

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return 13
}

func bar() (int, string) {
	return 11, "ok"
}
