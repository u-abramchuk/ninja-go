package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			runtime.Gosched()
		}
		wg.Done()
	}()

	go func() {
		for i := 20; i < 35; i++ {
			fmt.Println(i)
			runtime.Gosched()
		}
		wg.Done()
	}()

	wg.Wait()
}
