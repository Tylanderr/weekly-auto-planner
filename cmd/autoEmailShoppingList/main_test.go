package main

import "testing"

func TestParseIngredientString(t *testing.T) {

	item, count := parseIngredientString("2 apples")
	if count != 2 {
		t.Errorf("Expected 2, but got %d", count)
	}
	
	if item != "apple" {
		t.Errorf("Expected apple, but got %s", item)
	}
}
