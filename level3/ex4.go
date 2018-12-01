package main

import (
	"fmt"
)

func main() {
	year := 1987
	now := 2018
	for {
		if year > now {
			break
		}
		fmt.Println(year)
		year++
	}
}
