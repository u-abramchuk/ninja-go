package main

import (
	"fmt"
)

func main() {
	slc := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	slc = append(slc, 52)
	fmt.Println(slc)

	slc = append(slc, 53, 54, 55)
	fmt.Println(slc)

	x := []int{56, 57, 58, 59, 60}
	slc = append(slc, x...)
	fmt.Println(slc)
}
