package main

import (
	"fmt"
)

func main() {
	for n := 10; n <= 100; n++ {
		fmt.Println(n % 4)
	}
}
