package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0

	var wg sync.WaitGroup

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
