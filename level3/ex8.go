package main

import (
	"fmt"
)

func main() {
	switch {
	case 2 > 5:
		{
			fmt.Println("2 > 5")
		}
	case 5 == 5:
		{
			fmt.Println("5 == 5")
		}
	default:
		{
			fmt.Println("WTF!")
		}
	}
}
