package arrays

import "testing"

func TestSearchInArray(t *testing.T) {
	arrInt := [2]int{1, 2}
	arrString := [...]string{"Ram", "Shyam", "Rohan"}
	t.Run("int not in array", func(t *testing.T) {
		if _, err := linearSearch(arrInt, 3); err == nil {
			t.Errorf("Expected error got nil")
		}
	})

	t.Run("string not in array", func(t *testing.T) {
		if _, err := linearSearch(arrString, "mak"); err == nil {
			t.Errorf("Expected error got nil")
		}
	})

	t.Run("int in array", func(t *testing.T) {
		if _, err := linearSearch(arrInt, 2); err != nil {
			t.Errorf("Expected nil got error")
		}
	})

	t.Run("string in array", func(t *testing.T) {
		if _, err := linearSearch(arrString, "Shyam"); err != nil {
			t.Errorf("Expected nil got error")
		}
	})
}

func TestSearchInEmptyArray(t *testing.T) {
	arrInt := [5]int{}
	t.Run("empty array", func(t *testing.T) {
		if _, err := linearSearch(arrInt, 1); err == nil {
			t.Errorf("Expected error got nil")
		}
	})
	t.Run("empty array default value", func(t *testing.T) {
		if _, err := linearSearch(arrInt, 0); err != nil {
			t.Errorf("Expected error got nil")
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
