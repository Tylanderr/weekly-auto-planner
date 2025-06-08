package main

import (
	"reflect"
	"testing"

	"github.com/tylander732/weeklyAutoPlanner/pkg/model"
)

func TestSortIngredients(t *testing.T) {
	var ingredients = []model.Ingredient{
		{Name: "Chicken Breast", Count: 1},
		{Name: "Bread", Count: 1},
		{Name: "Chips", Count: 1},
		{Name: "Milk", Count: 1},
	}

	expectedIngredients := model.SortedIngredients{
		Produce:            make(map[string]int),
		MeatAndPoultry:     make(map[string]int),
		Seafood:            make(map[string]int),
		Dairy:              make(map[string]int),
		Bakery:             make(map[string]int),
		FrozenFoods:        make(map[string]int),
		PantryStaples:      make(map[string]int),
		Beverages:          make(map[string]int),
		Snacks:             make(map[string]int),
		HouseholdGoods:     make(map[string]int),
		PersonalCare:       make(map[string]int),
		InternationalFoods: make(map[string]int),
		Deli:               make(map[string]int),
		Floral:             make(map[string]int),
		Unsorted:           make(map[string]int),
	}

	expectedIngredients.MeatAndPoultry["Chicken Breast"] = 1
	expectedIngredients.Bakery["Bread"] = 1
	expectedIngredients.Snacks["Chips"] = 1
	expectedIngredients.Dairy["Milk"] = 1

	actualIngredients := model.SortedIngredients{
		Produce:            make(map[string]int),
		MeatAndPoultry:     make(map[string]int),
		Seafood:            make(map[string]int),
		Dairy:              make(map[string]int),
		Bakery:             make(map[string]int),
		FrozenFoods:        make(map[string]int),
		PantryStaples:      make(map[string]int),
		Beverages:          make(map[string]int),
		Snacks:             make(map[string]int),
		HouseholdGoods:     make(map[string]int),
		PersonalCare:       make(map[string]int),
		InternationalFoods: make(map[string]int),
		Deli:               make(map[string]int),
		Floral:             make(map[string]int),
		Unsorted:           make(map[string]int),
	}

	sortIngredients(ingredients, &actualIngredients)

	if !reflect.DeepEqual(actualIngredients, expectedIngredients) {
		t.Errorf("Structs did not match")
	}
}

func TestExecuteTemplate(t *testing.T) {

}
