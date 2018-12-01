package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Username string
	Age      int
}

func main() {
	users := []user{
		user{
			Username: "cool guy",
			Age:      32,
		},
		user{
			Username: "lone wolf",
			Age:      55,
		},
	}

	if serialized, err := json.Marshal(users); err == nil {
		fmt.Println(string(serialized))
	}
}
