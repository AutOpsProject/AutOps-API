package common

import (
	"testing"
)

type IntComparator struct{}

func (IntComparator) Compare(a, b int) int {
	// Standard integer comparison
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func TestListBasicOperations(t *testing.T) {
	list := NewList(IntComparator{}, []int{})

	// Test IsEmpty on empty list
	if !list.IsEmpty() {
		t.Error("expected list to be empty")
	}

	// Append values
	list.Append(3)
	list.Append(1)
	list.Append(2)
	list.Append(2)

	// Test GetItem
	index, val := list.GetItem(3)
	if index != 0 && *val != 3 {
		t.Errorf("expected index to be 0, got %d ; and val to be 3, got %d", index, *val)
	}
	index, val = list.GetItem(2)
	if index != 2 && *val != 2 {
		t.Errorf("expected index to be 2, got %d ; and val to be 2, got %d", index, *val)
	}
	index, val = list.GetItem(56)
	if index != -1 && val != nil {
		t.Errorf("expected index to be -1, got %d ; and val to be nil, got %p", index, val)
	}

	if list.Len() != 4 {
		t.Errorf("expected length 4, got %d", list.Len())
	}

	if list.IsEmpty() {
		t.Error("expected list to be non-empty after append")
	}

	// Test Contains
	if !list.Contains(1) {
		t.Error("expected list to contain 1")
	}
	if list.Contains(5) {
		t.Error("expected list to not contain 5")
	}
	if !list.Contains(2) {
		t.Error("expected list to contain 2")
	}

	// Test RemoveIndex
	err := list.RemoveIndex(1)
	if err != nil {
		t.Errorf("unexpected error in RemoveIndex: %v", err)
	}
	if list.Contains(1) {
		t.Error("expected list to not contain 1 after removal")
	}

	// Test Remove
	list.Remove(2)
	if !list.Contains(2) {
		t.Error("expected list to contain 2")
	}
	list.Remove(2)
	if list.Contains(2) {
		t.Error("expected list to not contain 2 after removal")
	}

	// Test RemoveIndex with invalid index
	err = list.RemoveIndex(10)
	if err == nil {
		t.Error("expected error for out-of-bounds index")
	}

	// Test RemoveAll
	list.Append(4)
	list.Append(3)
	list.RemoveAll(3)
	if list.Contains(3) {
		t.Error("expected all 3s to be removed")
	}

	// Test Sort
	list.Append(7)
	list.Append(2)
	list.Append(6)
	list.Sort()
	items := list.Items()
	for i := 1; i < len(items); i++ {
		if items[i-1] > items[i] {
			t.Errorf("list is not sorted: %v", items)
		}
	}

	// Test Clear
	list.Clear()
	if !list.IsEmpty() || list.Len() != 0 {
		t.Error("expected list to be empty after Clear")
	}
}

func TestSelectOne(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6}
	comparator := IntComparator{}
	list := List[int]{items, comparator}

	match, found := list.SelectOne(func(i int) bool {
		return i%97 == 0
	})
	if found {
		t.Errorf("expected to find no items x that matches x mod 97 == 0, but got %d", match)
	}

	match, found = list.SelectOne(func(i int) bool {
		return i%2 == 0
	})
	if !found {
		t.Error("expected to find an item that match x % 2 == 0, but got nothing")
	}
	if match != 2 {
		t.Errorf("expected 2, got %d", match)
	}
}

func TestSelectAll(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6}
	comparator := IntComparator{}
	list := List[int]{items, comparator}

	match, count := list.SelectAll(func(i int) bool {
		return i%2 == 0
	})

	if count != len(match) {
		t.Errorf("expected %d to be %d", count, len(match))
	} else if count != 3 {
		t.Errorf("expected %d to be 3", count)
	} else if match[0] != 2 || match[1] != 4 || match[2] != 6 {
		t.Errorf("expected %v to be [2, 4, 6]", match)
	}
}
