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
	IngredientsJArray []string `json:"ingredients"`
	SharedIngredients []string `json:"sharedIngredients"`
	Notes             string   `json:"notes"`
}

type Ingredients struct {
	Ingredient string `json:"ingredient"`
	Count      int    `json:"count"`
}

type EmailData struct {
	Receiver   string
	Meals      string
	Vegetables string
	Fruits     string
	Proteins   string
	Unsorted   string
}
