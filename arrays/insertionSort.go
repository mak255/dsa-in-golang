package arrays

import "fmt"

func insertionSort(arr []int) []int {
	fmt.Println("Running Insertion Sort")
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		for j := i - 1; j >= 0; j-- {
			if key < arr[j] {
				arr[j+1] = arr[j]
				arr[j] = key
				fmt.Println(arr)
			}
		}
	}
	fmt.Println("Finished Insertion Sort")
	return arr
}
