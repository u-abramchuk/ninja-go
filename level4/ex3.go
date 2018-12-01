package main

import (
	"fmt"
)

func main() {
	slc := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(slc[:5])
	fmt.Println(slc[5:])
	fmt.Println(slc[2:7])
	fmt.Println(slc[1:6])

}
