package main_test

import (
	"reflect"
	"testing"
)

type Item struct {
	ID   int
	Name string
}

func TestFindItem(t *testing.T) {
	items := []Item{
		{1, "one"},
		{2, "two"},
	}
	result1 := FindItem(items, func(item *Item) bool {
		return item.ID == 1
	})
	result2 := FindItem(items, func(item *Item) bool {
		return item.Name == "one"
	})
	if !reflect.DeepEqual(result1, result2) {
		t.Error("Unexpected results")
	}
}

// FindItem returns the item in the array that results from finder returning true.
func FindItem(items []Item, finder func(item *Item) bool) *Item {
	for i := range items {
		if finder(&items[i]) {
			return &items[i]
		}
	}
	return nil
}
