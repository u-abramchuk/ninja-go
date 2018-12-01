package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	for i := 0; i < 10; i++ {
		go func(k int) {
			for j := 0; j < 10; j++ {
				c <- k*10 + j
			}
		}(i)
	}

	for i := 0; i < 100; i++ {
		fmt.Println(<-c)
	}
}
