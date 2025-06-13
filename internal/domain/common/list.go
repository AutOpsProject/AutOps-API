package common

import "sort"

// Comparator defines a method to compare two values of type T.
//
// It should return -1 if a < b, 0 if a == b, and 1 if a > b.
type Comparator[T any] interface {
	Compare(a T, b T) int
}

// List represents a generic list of elements of type T, with operations based on a provided Comparator.
type List[T any] struct {
	items      []T
	comparator Comparator[T]
}

// NewList creates a new List with the given comparator and initial items.
func NewList[T any](comparator Comparator[T], items []T) *List[T] {
	return &List[T]{
		comparator: comparator,
		items:      items,
	}
}

// Append adds an item to the end of the list.
func (l *List[T]) Append(item T) {
	l.items = append(l.items, item)
}

// Remove deletes the first occurrence of the specified item from the list.
//
// It uses the comparator to determine equality.
func (l *List[T]) Remove(item T) {
	i := 0
	for i < len(l.items) && l.comparator.Compare(l.items[i], item) != 0 {
		i++
	}
	if i < len(l.items) {
		l.RemoveIndex(i)
	}
}

// RemoveIndex removes the element at the given index from the list.
//
// It returns an error if the index is out of bounds.
func (l *List[T]) RemoveIndex(index int) error {
	if index >= l.Len() {
		return ErrListIndexOutOfRange
	}
	l.items = append(l.items[:index], l.items[index+1:]...)
	return nil
}

// RemoveAll deletes all occurrences of the specified item from the list.
func (l *List[T]) RemoveAll(item T) {
	n := 0
	for _, elem := range l.items {
		if l.comparator.Compare(elem, item) != 0 {
			l.items[n] = elem
			n++
		}
	}
	l.items = l.items[:n]
}

// GetItem returns the index and a pointer to the first matching item in the list.
//
// If not found, it returns -1 and nil.
func (l *List[T]) GetItem(item T) (int, *T) {
	i := 0
	for i < len(l.items) && l.comparator.Compare(l.items[i], item) != 0 {
		i++
	}
	if i < len(l.items) {
		return i, &l.items[i]
	}
	return -1, nil
}

// Items returns a copy of the items in the list.
func (l *List[T]) Items() []T {
	return append([]T(nil), l.items...)
}

// Len returns the number of items in the list.
func (l *List[T]) Len() int {
	return len(l.items)
}

// Sort sorts the list using the comparator.
func (l *List[T]) Sort() {
	sort.Slice(l.items, func(i, j int) bool {
		return l.comparator.Compare(l.items[i], l.items[j]) < 0
	})
}

// Contains returns true if the item is present in the list.
func (l *List[T]) Contains(item T) bool {
	i := 0
	for i < len(l.items) && l.comparator.Compare(l.items[i], item) != 0 {
		i++
	}
	return i < len(l.items)
}

// Clear removes all items from the list.
func (l *List[T]) Clear() {
	l.items = l.items[:0]
}

// IsEmpty returns true if the list has no items.
func (l *List[T]) IsEmpty() bool {
	return l.Len() == 0
}

// SelectOne returns the first element in the list that satisfies the given filter function.
// If no element matches the filter, the zero value of T and false are returned.
//
// The filter function should return true for the element to be selected.
func (l *List[T]) SelectOne(filter func(T) bool) (T, bool) {
	for _, item := range l.items {
		if filter(item) {
			return item, true
		}
	}
	var zero T
	return zero, false
}

// SelectAll returns every element in the list that satisfies the given filter function, with the amount of matching items.
//
// The filter function should return true for the elements to be selected.
func (l *List[T]) SelectAll(filter func(T) bool) ([]T, int) {
	l2 := NewList(l.comparator, []T{})
	for _, item := range l.items {
		if filter(item) {
			l2.Append(item)
		}
	}
	return l2.Items(), len(l2.items)
}
