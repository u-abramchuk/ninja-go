package main

import (
	"fmt"
)

func main() {
	fmt.Println(struct {
		first string
		last  string
	}{
		first: "John",
		last:  "Bond",
	})
}
