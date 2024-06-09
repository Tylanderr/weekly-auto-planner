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

	"github.com/tylander732/autoEmailShoppingList/internal/consts"
	"github.com/tylander732/autoEmailShoppingList/internal/model"
	// "github.com/tylander732/autoEmailShoppingList/internal/projectpath"
)

var propertiesFile = "./resources/app.properties"
var username string
var password string

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

		//TODO: figure out what I'm passing as a parameter here
		for j := 0; j < len(meals); j++ {
			//for each meal, sort out the veggies, fruits and proteins
			sortedVegetables, sortedFruits, sortedProteins, unsorted := seperateIngredients(meals[j].IngredientsJArray)
			fmt.Println(sortedVegetables, sortedFruits, sortedProteins, unsorted)
			mealNames = append(mealNames, meals[j].Name) 
			fmt.Println(mealNames)
		}

		if err != nil {
			fmt.Println("Was unable to succesfully select meal for users", err)
		}

		// TODO: Now that I have the sorted lists of ingrediants, pass them into an 
		// html template for better formatting



		data := model.EmailData {
			Receiver: userArray[i].Email,
			Meals: mealNames,
		}

		executeTemplate("./resources/email_template.html", data)

		emailString := makeMealEmailString(meals)

		// sendEmail(emailString, emailReceivers)
	}
}

func readJsonFile() (model.JsonFile, error) {
	contents, err := os.ReadFile("./resources/userList.json")
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

// TODO: Rip out this function once I break things up into ingredient sections
func makeMealEmailString(meal []model.Meal) string {
	var emailString strings.Builder
	for i := 0; i < len(meal); i++ {
		currentMeal := meal[i]
		emailString.WriteString(currentMeal.Name + "\n")
		emailString.WriteString("Ingredients: ")
		emailString.WriteString(strings.Join(currentMeal.IngredientsJArray, ", "))
		emailString.WriteString("\n \n")
	}

	return emailString.String()
}

func sendEmail(emailString string, receivers []string) {
	// smtp server configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", username, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, username, receivers, []byte(emailString))
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

func seperateIngredients(ingredients []string) ([]string, []string, []string, []string) {

	localVegetables := []string{}
	localFruits := []string{}
	localProteins := []string{}
	localUnsorted := []string{}

	for i := 0; i < len(ingredients); i++ {
		currentIngredient := strings.ToLower(ingredients[i])

		if slices.Contains(consts.Vegetables, currentIngredient) {
			localVegetables = append(localVegetables, currentIngredient)
		} else if slices.Contains(consts.Fruits, currentIngredient) {
			localFruits = append(localFruits, currentIngredient)
		} else if slices.Contains(consts.Proteins, currentIngredient) {
			localProteins = append(localProteins, currentIngredient)
		} else {
			localUnsorted = append(localUnsorted, currentIngredient)
		}
	}

	return localVegetables, localFruits, localProteins, localUnsorted
}

func executeTemplate(templateFile string, data model.EmailData) (string, error) {
	// Parse the template file
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return "", err
	}

	// Execute the template with the provided data
	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, data)
	if err != nil {
		 return "", err
	}

	return tpl.String(), nil
}

//TODO: Make a function that will strip away the count of items needed when checking what category it will go into
// Example: 5x eggs - will simplify down to just "eggs" when checking categories
