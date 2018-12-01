package main

import (
	"fmt"
)

type customErr struct {
	details string
}

func (e customErr) Error() string {
	return fmt.Sprintf("Custom error occured. Details: %v", e.details)
}

func foo(e error) {
	fmt.Println("Foo!", e)
}

func main() {
	e := customErr{
		details: "Something bad has happened",
	}

	foo(e)
}
