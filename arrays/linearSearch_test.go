package arrays

import (
	"errors"
	"testing"
)

func TestSearchInSlice(t *testing.T) {
	arrInt := []int{1, 2}
	arrString := []string{"Ram", "Shyam", "Rohan"}
	t.Run("int not in array", func(t *testing.T) {
		_, err := linearSearch(arrInt, 3)
		if !(errors.Is(err, ItemNotFound)) {
			t.Errorf("Expected %v got %v", ItemNotFound, err)
		}
	})

	t.Run("string not in array", func(t *testing.T) {
		_, err := linearSearch(arrString, "mak")
		if !(errors.Is(err, ItemNotFound)) {
			t.Errorf("Expected %v got %v", ItemNotFound, err)
		}
	})

	t.Run("int in array", func(t *testing.T) {
		i, err := linearSearch(arrInt, 2)
		if err != nil {
			t.Errorf("Expected %v got %v", nil, err)
		}
		if arrInt[i] != 2 {
			t.Errorf("Expected 2 got %v", arrInt[i])
		}
	})

	t.Run("string in array", func(t *testing.T) {
		i, err := linearSearch(arrString, "Shyam")
		if err != nil {
			t.Errorf("Expected %v got %v", nil, err)
		}
		if arrString[i] != "Shyam" {
			t.Errorf("Expected Shyam got %v", arrString[i])
		}
	})
}

func TestSearchInEmptySlice(t *testing.T) {
	arrInt := []int{}
	t.Run("empty array default value", func(t *testing.T) {
		_, err := linearSearch(arrInt, 0)
		if !(errors.Is(err, EmptySlice)) {
			t.Errorf("Expected %v, got %v", EmptySlice, err)
		}
	})
}

func TestSort(t *testing.T) {
	arr := [5]int{2, 11, 4, 9, 8}
	sortedArr := [5]int{2, 4, 8, 9, 11}
	a := selectionSort(arr)
	if a != sortedArr {
		t.Errorf("Expected %v got this %v", sortedArr, a)
	}
	a = bubbleSort(arr)
	if a != sortedArr {
		t.Errorf("Expected %v got this %v", sortedArr, a)
	}
}
