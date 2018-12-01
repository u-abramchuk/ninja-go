package main

import (
	"fmt"
)

func main() {
	items := []int{1, 4, 7, 2, 6, 4}

	fmt.Println(foo(items...))
	fmt.Println(bar(items))
}

func foo(xi ...int) int {
	result := 0

	for _, v := range xi {
		result += v
	}

	return result
}

func bar(xi []int) int {
	result := 0

	for _, v := range xi {
		result += v
	}

	return result
}
