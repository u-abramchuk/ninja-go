package main

import (
	"fmt"
)

func main() {
	slc := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	y := append(slc[:3], slc[6:]...)
	fmt.Println(y)
}
