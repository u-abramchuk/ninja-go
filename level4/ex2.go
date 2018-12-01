package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 43, 6, 4, 6, 143}
	fmt.Println(arr)

	for i, v := range arr {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", arr)
}
