package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"math/rand"

	"github.com/magiconair/properties"

	"net/smtp"
	"os"
	"slices"
	"strings"

	"github.com/tylander732/autoEmailShoppingList/pkg/consts"
	"github.com/tylander732/autoEmailShoppingList/pkg/model"
)

var propertiesFile = "./resources/app.properties"
var username string
var password string

var groceryCategories = []consts.GroceryCategory{
	consts.Produce,
	consts.MeatAndPoultry,
	consts.Seafood,
	consts.Dairy,
	consts.Bakery,
	consts.FrozenFoods,
	consts.PantryStaples,
	consts.Beverages,
	consts.Snacks,
	consts.HouseholdGoods,
	consts.PersonalCare,
	consts.InternationalFoods,
	consts.Deli,
	consts.Floral,
}

func main() {
	readProperties()
	userJsonFile, errReadingJson := readJsonFile()
	if errReadingJson != nil {
		fmt.Println("Error gathering file from disk: ", errReadingJson)
		return
	}

	userArray := userJsonFile.UserJArray

	for i := 0; i < len(userArray); i++ {
		meals, err := selectMeals(userArray[i])

		mealNames := []string{}

		// For each user, initialize a sortedIngredientsStruct
		sortedIngredientsStruct := model.SortedIngredients{
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

		for j := 0; j < len(meals); j++ {
			mealNames = append(mealNames, meals[j].Name)

			sortIngredients(meals[j].Ingredients, &sortedIngredientsStruct)
		}

		if err != nil {
			fmt.Println("Was unable to succesfully select meal for users", err)
		}

		mealsString := strings.Join(mealNames, "\n")
		mealsString += "\n"

		data := model.EmailData{
			Receiver: userArray[i].Email,
			Meals:    mealsString,
		}

		emailBody, err := executeTemplate("./resources/email_template.html", data)

		sendEmail(emailBody, userArray[i].Email)
	}
}

func readJsonFile() (model.JsonFile, error) {
	contents, err := os.ReadFile("./resources/userListUpdated.json")
	if err != nil {
		return model.JsonFile{}, err
	}

	data := model.JsonFile{}

	err = json.Unmarshal(contents, &data)

	if err != nil {
		fmt.Println(err)
	}

	return data, nil

}

func selectMeals(usersData model.User) ([]model.Meal, error) {
	numOfUsersMeal := len(usersData.MealJArray)
	numOfmealsToSelect := usersData.NumOfMealsToSelect

	mealsToSend := []model.Meal{}

	randomMealsToBeSelected, err := generateUniqueRandomIntegers(numOfUsersMeal, numOfmealsToSelect)
	if err != nil {
		return []model.Meal{}, err
	}

	for i := 0; i < len(randomMealsToBeSelected); i++ {
		mealsToSend = append(mealsToSend, usersData.MealJArray[randomMealsToBeSelected[i]])
	}

	//get the meal objects
	return mealsToSend, nil
}

func generateUniqueRandomIntegers(numberRange int, amountToGenerate int) ([]int, error) {
	if amountToGenerate > numberRange {
		return []int{}, errors.New("Amount asked to generate, is higher than the number of user meals added")
	}
	uniqueInts := []int{}

	i := 0
	// I might have an off by one error in this loop, we'll see
	for i < amountToGenerate {
		currentInt := rand.Intn(numberRange)
		if !slices.Contains(uniqueInts, currentInt) {
			uniqueInts = append(uniqueInts, currentInt)
			i++
		}
	}

	return uniqueInts, nil
}

func sendEmail(emailString string, receiver string) {
	// smtp server configuration
	smtpHost := "smtp.gmail.com"
	server := smtpHost + ":587"

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	auth := smtp.PlainAuth("", username, password, smtpHost)

	subject := "Subject: Auto Emailer Test\n"

	msg := []byte(subject + mime + emailString)

	err := smtp.SendMail(server, auth, username, []string{receiver}, msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email sent successfully")
}

func readProperties() {
	p := properties.MustLoadFile(propertiesFile, properties.UTF8)
	username, _ = p.Get("username")
	password, _ = p.Get("password")
}

// Receives a list of ingredients from a model.Meal
// Checks current ingredient against grocery types defined in pkg/consts/consts.go
// Returns struct containing sortedIngredients and running count for each item
func sortIngredients(ingredients []model.Ingredient, sortedIngredients *model.SortedIngredients) {

	// For each ingredient, check if it is contained within one of the slices
	for i := 0; i < len(ingredients); i++ {
		ci := ingredients[i]

		for _, slice := range groceryCategories {
			// If found within current grocery category
			if stringInSlice(ci.Name, slice.ItemsSlice) {
				sortedIngredients.IncrementIngredientCount(slice.Name, ci)
				break
			}
		}
	}
}

// Check if a given string is within a slice
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func executeTemplate(templateFile string, data model.EmailData) (string, error) {

	// Parse the template file
	template, err := template.ParseFiles(templateFile)
	if err != nil {
		return "", err
	}

	// Execute the template with the provided data
	var tpl bytes.Buffer
	err = template.Execute(&tpl, data)
	if err != nil {
		fmt.Print("There has been an error executing the HTML Template: ")
		fmt.Println(err)
		return "", err
	}

	return tpl.String(), nil
}
