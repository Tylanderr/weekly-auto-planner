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
	Name              string       `json:"name"`
	Ingredients       []Ingredient `json:"ingredients"`
	// SharedIngredients []string     `json:"sharedIngredients"`
	Notes             string       `json:"notes"`
}

type Ingredient struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type EmailData struct {
	Receiver       string
	Meals          string
	AllIngredients SortedIngredients
}

type SortedIngredients struct {
	Produce            map[string]int
	MeatAndPoultry     map[string]int
	Seafood            map[string]int
	Dairy              map[string]int
	Bakery             map[string]int
	FrozenFoods        map[string]int
	PantryStaples      map[string]int
	Beverages          map[string]int
	Snacks             map[string]int
	HouseholdGoods     map[string]int
	PersonalCare       map[string]int
	InternationalFoods map[string]int
	Deli               map[string]int
	Floral             map[string]int
	Unsorted           map[string]int
}

func (s *SortedIngredients) IncrementIngredientCount(category string, ingredient Ingredient) {
	var targetMap map[string]int

	switch category {
	case "Produce":
		targetMap = s.Produce
	case "MeatAndPoultry":
		targetMap = s.MeatAndPoultry
	case "Seafood":
		targetMap = s.Seafood
	case "Dairy":
		targetMap = s.Dairy
	case "Bakery":
		targetMap = s.Bakery
	case "FrozenFoods":
		targetMap = s.FrozenFoods
	case "PantryStaples":
		targetMap = s.PantryStaples
	case "Beverages":
		targetMap = s.Beverages
	case "Snacks":
		targetMap = s.Snacks
	case "HouseholdGoods":
		targetMap = s.HouseholdGoods
	case "PersonalCare":
		targetMap = s.PersonalCare
	case "InternationalFoods":
		targetMap = s.InternationalFoods
	case "Deli":
		targetMap = s.Deli
	case "Floral":
		targetMap = s.Floral
	default:
		targetMap = s.Unsorted
	}

	// Increment the count in the target map:
	targetMap[ingredient.Name] = targetMap[ingredient.Name] + ingredient.Count

}
