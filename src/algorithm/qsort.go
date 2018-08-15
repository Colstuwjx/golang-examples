package main

import (
	"log"
)

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func partition(data []int, left, right, pivotIndex int) (newPivotIndex int) {
	pivotValue := data[pivotIndex]
	swap(data, pivotIndex, right)
	storedIndex := left

	i := left
	for i <= right-1 {
		if data[i] < pivotValue {
			swap(data, i, storedIndex)
			storedIndex = storedIndex + 1
		}

		i++
	}

	swap(data, right, storedIndex)
	return storedIndex
}

func quickSort(data []int, left, right int) {

	if right > left {
		pivotIndex := left
		pivotNewIndex := partition(data, left, right, pivotIndex)
		log.Println("Pivot: ", data[pivotNewIndex], ", Data: ", data, ", Pivot index: ", pivotNewIndex)
		log.Println("Left: ", left, "Right: ", right, "Pivot: ", pivotNewIndex)

		quickSort(data, left, pivotNewIndex)
		quickSort(data, pivotNewIndex+1, right)
	}
}

func main() {
	data := []int{32, 4, 2, 1, 5, 323, 24, 53, 5, 4, 3, 63, 3, 123}
	quickSort(data, 0, len(data)-1)
}
