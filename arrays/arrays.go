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
