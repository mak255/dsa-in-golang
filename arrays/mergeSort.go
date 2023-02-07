package arrays

import (
	"fmt"
)

func mergeSort(arr []int) {
	fmt.Println("Starting merge sort")
	fmt.Println(arr)
	if len(arr) > 1 {
		mid := len(arr) / 2
		lArray := arr[:mid]
		rArray := arr[mid:]
		mergeSort(lArray)
		mergeSort(rArray)
		tmp := merge(lArray, rArray)
		for i := range tmp {
			arr[i] = tmp[i]
		}
	}
	fmt.Println("Ended merge sort")
}

func merge(arr1, arr2 []int) []int {
	fmt.Println("Starting Merge")
	var i, j, k int
	var tmp = make([]int, len(arr1)+len(arr2))
	fmt.Printf("arr1: %v \n", arr1)
	fmt.Printf("arr2: %v \n", arr2)
	for (i < len(arr1)) && (j < len(arr2)) {
		if arr1[i] < arr2[j] {
			tmp[k] = arr1[i]
			i++
		} else {
			tmp[k] = arr2[j]
			j++
		}
		k++
	}
	for i < len(arr1) {
		tmp[k] = arr1[i]
		k++
		i++
	}
	for j < len(arr2) {
		tmp[k] = arr2[j]
		j++
		k++
	}
	fmt.Println(tmp)
	return tmp
}
