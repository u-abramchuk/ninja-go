package main

import (
	"fmt"
)

func main() {
	fmt.Println("Anon func says", func() int {
		return 13
	}())
}
