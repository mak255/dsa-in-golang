package arrays

import (
	"errors"
	"fmt"
)

var ItemNotFound = errors.New("Item not present in array")
var EmptySlice = errors.New("Trying to search in empty slice")

func linearSearch[T int | float64 | float32 | string](arr []T, value T) (int, error) {
	if len(arr) == 0 {
		return 0, EmptySlice
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == value {
			return i, nil
		}
	}
	return 0, ItemNotFound
}

func selectionSort(arr [5]int) [5]int {
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]

		}
		fmt.Println(arr)
	}
	return arr
}

func bubbleSort(arr [5]int) [5]int {
	for i := 0; i < len(arr); i++ {
		swapped := false
		for j := 0; j < len(arr)-(i+1); j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				swapped = true
			}
		}
		if !swapped {
			break
		}
		fmt.Println(arr)
	}
	return arr
}
