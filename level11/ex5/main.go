package main

import (
	"fmt"
)

func main() {
	fmt.Println(Sum(2, 2))
}

// Sum adds two integers
func Sum(a int, b int) int {
	return a + b
}
