package main

import (
	"fmt"
)

func main() {
	var kvMap = map[string]string{
		"A": "1",
		"B": "2",
	}

	if _, exists := kvMap["foo"]; exists {
		fmt.Println("there is key foo in map.")
	} else {
		fmt.Println("foo doesn't exist in map.")
	}
}
