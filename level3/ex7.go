package main

import (
	"fmt"
)

func main() {
	if 2 > 5 {
		fmt.Println("2 > 5")
	} else if 5 == 5 {
		fmt.Println("5 == 5")
	} else {
		fmt.Println("WTF!")
	}
}
