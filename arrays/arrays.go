package arrays

import (
	"fmt"
	"reflect"
)

func linearSearch(arr interface{}, value interface{}) (int, error) {
	a := reflect.ValueOf(arr)

	if a.Kind() != reflect.Array {
		return 0, fmt.Errorf("Not an Array")
	}

	for i := 0; i < a.Len(); i++ {
		if a.Index(i).Interface() == value {
			return i, nil
		}
	}
	return 0, fmt.Errorf("Item not found in Array")
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
