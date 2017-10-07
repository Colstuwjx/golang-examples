package main

/*
   anonymous function makes code reuseable!
*/

import (
	"fmt"
)

func filter(src []string, fn func(string) bool) []string {
	var filtered []string

	for _, item := range src {
		if fn(item) == true {
			filtered = append(filtered, item)
		}
	}

	return filtered
}

func filterONLYJacky(item string) bool {
	if item == "Jacky" {
		return true
	} else {
		return false
	}
}

func main() {
	var src []string = []string{
		"Jacky",
		"Tom",
		"John",
		"Carrie",
		"Cathy",
	}

	filtered := filter(src, filterONLYJacky(item))
	fmt.Println("Filtered values: ", filtered)
}
