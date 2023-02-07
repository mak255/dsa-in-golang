package arrays

import (
	"reflect"
	"testing"
)

func TestSorting(t *testing.T) {
	unsortedslice := []int{23, 35, 1, 5, 3, 7, 4, 2, 88}
	sortedSlice := []int{1, 2, 3, 4, 5, 7, 23, 35, 88}
	t.Run("InsertionSort", func(t *testing.T) {
		if !(reflect.DeepEqual(sortedSlice, insertionSort(unsortedslice))) {
			t.Errorf("Expected %v got %v", sortedSlice, unsortedslice)
		}
	})
	t.Run("MergeSort", func(t *testing.T) {
		mergeSort([]int{23, 35, 1, 5, 3, 7, 4, 2, 88})
		if !(reflect.DeepEqual(sortedSlice, sortedSlice)) {
			t.Errorf("Expected %v got %v", sortedSlice, unsortedslice)
		}
	})
}
