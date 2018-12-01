package main

import (
	"fmt"
)

const (
	year2018 = 2018 - iota
	year2017 = 2018 - iota
	year2016 = 2018 - iota
	year2015 = 2018 - iota
)

func main() {
	fmt.Println(year2018)
	fmt.Println(year2017)
	fmt.Println(year2016)
	fmt.Println(year2015)
}
