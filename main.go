package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "math/rand"
    "net/smtp"
    "time"
)

type JsonFile struct {
    UserJArray []User `json:"users"`
}

type User struct {
    Email string `json:"email"`
    MealJArray []Meal `json:"meals"`
}

type Meal struct {
    Name string `json:"name"`
    IngrediantsJArray []string `json:"ingrediants"`
}

func main() {
    rand.Seed(time.Now().UnixNano())
    userData, errReadingJson := readJsonFile()
    if errReadingJson != nil {
        fmt.Println("Error gathering file from disk: ", errReadingJson)
        return
    }

    selectMeal(userData)

}

func readJsonFile() (JsonFile, error) {
    contents, err := ioutil.ReadFile("./resources/userList.json")
    if err != nil {
        return JsonFile{}, err
    }
    // fmt.Println(string(contents))

    data := JsonFile{}

    err = json.Unmarshal(contents, &data)

    if err != nil {
        fmt.Println(err)
    }

    return data, nil

}

func selectMeal(userData JsonFile) {
    for i := 0; i < len(userData.UserJArray); i++ {
        // for the current user, see how many meals they have added
        // select 3 random numbers in that range
        // prepare email to be sent with those 3 meals and their ingrediants
        fmt.Println(userData.UserJArray[i])
        var numOfUsersMeal = len(userData.UserJArray[i].MealJArray)
        meal1, meal2, meal3 := generateRandomIntegers(numOfUsersMeal)
        fmt.Println(meal1, meal2, meal3)
    }
}

func generateRandomIntegers(numberRange int) (int, int, int) {
    // we need these to be unique numbers
    return rand.Intn(numberRange + 1), rand.Intn(numberRange + 1), rand.Intn(numberRange + 1)
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


