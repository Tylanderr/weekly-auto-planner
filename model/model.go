package model


type JsonFile struct {
	UserJArray []User `json:"users"`
}

type User struct {
	Email              string `json:"email"`
	NumOfMealsToSelect int    `json:"numOfmealsToSelect"`
	MealJArray         []Meal `json:"meals"`
}

type Meal struct {
	Name              string   `json:"name"`
	IngrediantsJArray []string `json:"ingrediants"`
}
