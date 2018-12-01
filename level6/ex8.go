package main

import (
	"fmt"
)

func main() {
	even := isDivisibleBy(2)

	fmt.Println("4 is even", even(4))
}

func isDivisibleBy(n int) func(int) bool {
	return func(num int) bool {
		return num%n == 0
	}
}
