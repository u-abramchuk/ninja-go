package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type users []user

func (u users) Len() int {
	return len(u)
}

func (u users) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

type ByLast struct {
	users
}

func (u ByLast) Less(i, j int) bool {
	return u.users[i].Last < u.users[j].Last
}

type ByAge struct {
	users
}

func (u ByAge) Less(i, j int) bool {
	return u.users[i].Age < u.users[j].Age
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	print(users)

	sort.Sort(ByLast{users})

	print(users)

	sort.Sort(ByAge{users})

	print(users)
}

func print(users []user) {
	fmt.Println("===============")

	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)

		for _, s := range u.Sayings {
			fmt.Println("\t", s)
		}
	}
}
