package main

import (
	"fmt"
	"testing"
)

func TestParseIngredientString(t *testing.T) {

	parsedItem, count, err := parseIngredientString("2 apples")
	// __AUTO_GENERATED_PRINT_VAR_START__
	fmt.Println(fmt.Sprintf("TestParseIngredientString count: %v", count)) // __AUTO_GENERATED_PRINT_VAR_END__
	if err != nil {
		fmt.Println("Error returned from parseIngredientString")
	}
	if count != 2 {
		t.Errorf("Expected 2, but got %d", count)
	}
	
	if parsedItem != "apple" {
		t.Errorf("Expected apple, but got %s", parsedItem)
	}
}
