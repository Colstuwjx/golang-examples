package main

import (
	"fmt"
)

// map, slice, channel, both them have a pointer which point to the real value.
// map struct {
//     m *internalHashtable
// }

// channel struct {
//     c *internalChannel
// }

// slice struct {
//     array *internalArray
//     len   int
//     cap   int
// }
// thus, these type would be `reference type`!

func changeArrayValue(testValue [6]int) {
	testValue[0] = 1
}

func changeSliceValue(testValue []int) {
	testValue[0] = 1
}

func main() {
	primeArray := [6]int{2, 3, 5, 7, 11, 13}
	primeSlice := []int{2, 3, 5, 7, 11, 13}

	changeArrayValue(primeArray)
	changeSliceValue(primeSlice)

	// the first item will still be `2`
	fmt.Println(primeArray)

	// the first item will be `1`
	fmt.Println(primeSlice)
}
