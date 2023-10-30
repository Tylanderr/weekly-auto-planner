package main

import (
    "encoding/json"
    "fmt"
    "os"
    "math/rand"
    "net/smtp"
    "slices"
)

type JsonFile struct {
    UserJArray []User `json:"users"`
}

type User struct {
    Email string `json:"email"`
    NumOfMealsToSelect int `json:"numOfmealsToSelect"`
    MealJArray []Meal `json:"meals"`
}

type Meal struct {
    Name string `json:"name"`
    IngrediantsJArray []string `json:"ingrediants"`
}

func main() {
    userData, errReadingJson := readJsonFile()
    if errReadingJson != nil {
        fmt.Println("Error gathering file from disk: ", errReadingJson)
        return
    }

    selectMeal(userData)
}

func readJsonFile() (JsonFile, error) {
    contents, err := os.ReadFile("./resources/userList.json")
    if err != nil {
        return JsonFile{}, err
    }

    data := JsonFile{}

    err = json.Unmarshal(contents, &data)

    if err != nil {
        fmt.Println(err)
    }

    return data, nil

}

func selectMeal(userData JsonFile) {
    for i := 0; i < len(userData.UserJArray); i++ {
        var numOfUsersMeal = len(userData.UserJArray[i].MealJArray)
        numOfmealsToSelect := userData.UserJArray[i].NumOfMealsToSelect
        if numOfmealsToSelect < numOfUsersMeal {
            fmt.Println("User had a numOfmealsToSelect higher than the number of meals added")
            return
        }
        fmt.Println("About to generate numbers")
        meal1, meal2, meal3 := generateUniqueRandomIntegers(numOfUsersMeal, numOfmealsToSelect)
        fmt.Println(meal1, meal2, meal3)
    }
}

func generateUniqueRandomIntegers(numberRange int, amountToGenerate int) (int, int, int) {
    uniqueInts := []int{}

    i := 0
    for i < 3 {
        currentInt := rand.Intn(numberRange + 1)
        if !slices.Contains(uniqueInts, currentInt) {
            uniqueInts = append(uniqueInts, currentInt)
            i++
        }
    }

    return uniqueInts[0], uniqueInts[1], uniqueInts[2]
}


func sendEmail() {
    from := ""
    password := ""

    to := []string{""}

    // smtp server configuration
    smtpHost := "smtp.gmail.com"
    smtpPort := "587"
    message := []byte("This is a test email message")

    auth := smtp.PlainAuth("", from, password, smtpHost)

    err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Email sent successfully")
}


