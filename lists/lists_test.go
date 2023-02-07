package lists

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	l := NewList()
	if l.length != 0 {
		t.Errorf("Expected 0 got %v", l.length)
	}
	if l.head != nil {
		t.Errorf("Expected nil got %v", l.head)
	}
}

func TestPushinList(t *testing.T) {
	itemsToPush := []int{1, 2, 3}

	l := NewList()
	t.Run("Push to LinkedList", func(t *testing.T) {
		for i := range itemsToPush {
			l.Push(itemsToPush[i])
			if l.length != i+1 {
				t.Errorf("Expected %d got %d", i+1, l.length)
			}
			if l.head.data != itemsToPush[i] {
				t.Errorf("Expected %d got %d", itemsToPush[i], l.head.data)
			}
		}
	})
	t.Run("2nd item in LinkedList", func(t *testing.T) {
		if l.head.next.data != 2 {
			t.Errorf("Expected 2 got %v", l.head.next.data)
		}
	})
	l.Println()

}

func TestAppendInList(t *testing.T) {
	itemsToAppend := []int{1, 2, 3}
	l := NewList()
	t.Run("Append to LinkedList", func(t *testing.T) {
		for i := range itemsToAppend {
			l.Append(itemsToAppend[i])
			if l.length != i+1 {
				t.Errorf("Expected %d got %d", i+1, l.length)
			}
		}
	})
	t.Run("1st item", func(t *testing.T) {
		if l.head.data != 1 {
			t.Errorf("Expected %d got %d", l.head.data, itemsToAppend[0])
		}
	})
	l.Println()
}
