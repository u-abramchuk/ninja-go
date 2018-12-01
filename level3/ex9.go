package main

import (
	"fmt"
)

func main() {
	favSport := "Racing"

	switch favSport {
	case "Racing":
		fmt.Println("Racing")
	case "Football":
		fmt.Println("Football")
	default:
		fmt.Println("WTF!")
	}
}
