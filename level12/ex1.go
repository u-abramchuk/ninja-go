package main

import (
	"fmt"
	"./dog"
)

func main(){
	var years int

	fmt.Println("Years: ")
	
	fmt.Scan(&years)

	fmt.Printf("That's %v dog years", dog.Years(years))
}