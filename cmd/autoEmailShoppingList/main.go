package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"math/rand"
	"strconv"

	"github.com/magiconair/properties"

	"net/smtp"
	"os"
	"slices"
	"strings"

	"github.com/tylander732/autoEmailShoppingList/pkg/consts"
	"github.com/tylander732/autoEmailShoppingList/pkg/model"
	// "github.com/tylander732/autoEmailShoppingList/pkg/projectpath"
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

		var (
			mealNames        = []string{}
			sortedVegetables = []string{}
			sortedFruits     = []string{}
			sortedProteins   = []string{}
			unsorted         = []string{}
		)

		for j := 0; j < len(meals); j++ {
			mealNames = append(mealNames, meals[j].Name)
			fmt.Println(mealNames)

			sortedVegetables, sortedFruits, sortedProteins, unsorted = seperateIngredients(meals[j].IngredientsJArray)
		}

		if err != nil {
			fmt.Println("Was unable to succesfully select meal for users", err)
		}

		mealsString := strings.Join(mealNames, "\n")
		mealsString += "\n"

		veggiesString := strings.Join(sortedVegetables, ", ")
		fruitsString := strings.Join(sortedFruits, ", ")
		proteinsString := strings.Join(sortedProteins, ", ")
		unsortedString := strings.Join(unsorted, ", ")

		data := model.EmailData{
			Receiver:   userArray[i].Email,
			Meals:      mealsString,
			Vegetables: veggiesString,
			Fruits:     fruitsString,
			Proteins:   proteinsString,
			Unsorted:   unsortedString,
		}

		emailBody, err := executeTemplate("./resources/email_template.html", data)
		// __AUTO_GENERATED_PRINT_VAR_START__
		fmt.Println(fmt.Sprintf("main emailBody: %v", emailBody)) // __AUTO_GENERATED_PRINT_VAR_END__

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

//TODO: Fix. This is going to be broken after ingredient object introduction
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
