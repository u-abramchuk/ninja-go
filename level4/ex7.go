package main

import (
	"fmt"
)

func main() {
	slice := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}
	for _, v := range slice {
		for _, val := range v {
			fmt.Print(val, "\t")
		}
		fmt.Printf("\n")
	}
}
