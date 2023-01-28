package arrays

import "testing"

func TestItemNotInArray(t *testing.T) {
	arrInt := [2]int{1, 2}
	arrString := [...]string{"Ram", "Shyam", "Rohan"}
	_, err := linearSearch(arrInt, 3)
	if err == nil {
		t.Error("Wrong result")
	}
	_, err = linearSearch(arrString, "Hello")
	if err == nil {
		t.Error("Wrong result")
	}
}

func TestItemInArray(t *testing.T) {
	arrInt := [3]int{1, 2, 3}
	arrString := [...]string{"Ram", "Shyam", "Rohan"}
	_, err := linearSearch(arrInt, 1)
	if err != nil {
		t.Error("Wrong result")
	}
	_, err = linearSearch(arrString, "Ram")
	if err != nil {
		t.Error("Wrong result")
	}
}
